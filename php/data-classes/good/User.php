<?php

class User
{
    private $id;
    private $name;
    private $email;
    private $age;
    private $emailValidator;

    public function __construct($id, $name, $email, $age, EmailValidator $emailValidator = null)
    {
        $this->id = $id;
        $this->name = $name;
        $this->age = $age;
        $this->emailValidator = $emailValidator ?: new EmailValidator();

        $this->setEmail($email); // Use setter for validation
    }

    public function getId()
    {
        return $this->id;
    }

    public function getName()
    {
        return $this->name;
    }

    public function getEmail()
    {
        return $this->email;
    }

    public function setEmail($email)
    {
        if (!$this->emailValidator->isValid($email)) {
            throw new InvalidArgumentException('Invalid email address');
        }
        $this->email = $email;
    }

    public function getAge()
    {
        return $this->age;
    }

    public function setAge($age)
    {
        if ($age < 0 || $age > 150) {
            throw new InvalidArgumentException('Age must be between 0 and 150');
        }
        $this->age = $age;
    }

    public function getDisplayName()
    {
        return $this->name . ' (' . $this->age . ' years old)';
    }

    public function canVote()
    {
        return $this->age >= 18;
    }

    public function isAdult()
    {
        return $this->age >= 18;
    }

    public function getAgeCategory()
    {
        if ($this->age < 13) return 'child';
        if ($this->age < 20) return 'teenager';
        if ($this->age < 65) return 'adult';
        return 'senior';
    }
}

class EmailValidator
{
    public function isValid($email)
    {
        return filter_var($email, FILTER_VALIDATE_EMAIL) !== false;
    }
}
