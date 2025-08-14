# Password Strength Checker

This is a simple password strength checker written in Go. It uses common criteria such as length, capital letters, numbers, etc.
Uses HaveIBenPwned API to check if the password has been compromised, by sending the first characters of the password hash and checking locally for matches so no passwords of yours get leaked.

Run it and type in your password. You will get an evaluation based on said criteria, including feedback on how you could improve.
If your password has been compromised, it will let you know at the end.

[Installation](#installation-on-linux-based-systems)
[Usage](#usage)

## Installation on Linux-based systems
```bash
git clone https://github.com/raymiamis/passwordstrength.git
cd passwordstrength
./install.sh
```

## Installation on Windows or Mac
(Will require Go as a dependency)
```
go install https://github.com/raymiamis/passwordstrength.git
```
## Usage
```
passwordstrength
```
And then fill out your password you want to check. You will then get our evaluation.
