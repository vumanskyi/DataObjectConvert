# Data Object Converter

The CLI tool to turn JSON, YAML or XML objects into Data Object (PHP)

## Usage

Convert **json** structure to php data object 
```
data-object-convert json  -c PersonDto "{\"name\": \"Vlad\", \"age\": 30, \"cost\": 30.00}"
```

Result:
```
<?php

declare(strict_types=1);

class PerstonDto
{
    private int $cost;

    private string $name;

    private int $age;

    public function __construct(
        int $cost,
        string $name,
        int $age
    ) {
        $this->name = name;
        $this->age = age;
        $this->cost = cost;
    }

    public function getAge(): int
    {
        return $this->age;
    }

    public function getCost(): int
    {
        return $this->cost;
    }

    public function getName(): string
    {
        return $this->name;
    }

}

```

Convert **yaml** structure to php data object

```
data-object-convert yaml -c PersonDto "
name: Vlad
age: 30
cost: 30
"
```