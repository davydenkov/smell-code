<?php

class UserManager {
    private UserValidator $validator;
    private UserRepository $repository;
    private EmailService $emailService;
    private NotificationService $notificationService;

    public function __construct($db) {
        $this->validator = new UserValidator();
        $this->repository = new UserRepository($db);
        $this->emailService = new EmailService();
        $this->notificationService = new NotificationService($db);
    }

    public function registerUser(array $userData): int {
        $this->validator->validateRegistrationData($userData);

        if ($this->repository->userExists($userData['email'])) {
            throw new Exception('User already exists');
        }

        $userData = $this->prepareUserData($userData);

        $userId = $this->repository->createUser($userData);
        $this->repository->createUserProfile($userId, $userData);
        $this->repository->createUserSettings($userId);

        $this->emailService->sendVerificationEmail(
            $userData['email'],
            $userData['firstName'],
            $userData['verificationToken']
        );

        $this->notificationService->sendWelcomeNotification($userId);
        Logger::logRegistration($userData['email']);

        return $userId;
    }

    private function prepareUserData(array $userData): array {
        $userData['password'] = password_hash($userData['password'], PASSWORD_DEFAULT);
        $userData['verificationToken'] = bin2hex(random_bytes(32));
        return $userData;
    }
}
