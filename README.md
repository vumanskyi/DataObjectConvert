# Data Object Converter

The CLI tool to turn JSON and YAML objects into Data Object (PHP)

There are two types of object: entity and data transfer object.

## Usage

Convert **json** structure to php data object 
```
data-object-convert json  -c PersonDto "{\"name\": \"John\", \"age\": 30, \"cost\": 30.20}"
```

Result:
```
<?php

declare(strict_types=1);

class PerstonDto
{
    private float $cost;

    private string $name;

    private int $age;

    public function __construct(
        float $cost,
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

    public function getCost(): float
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
data-object-convert yaml -t entity -c PersonDto "
name: John
age: 30
cost: 30.20
"
```

Result:

```
<?php

declare(strict_types=1);

class PersonDto
{
    private string $name;

    private int $age;

    private float $cost;

    public function setName(string $name): self
    {
        return $this->name = $name;
        return $this;
    }

    public function getName(): string
    {
        return $this->name;
    }

    public function setAge(int $age): self
    {
        return $this->age = $age;
        return $this;
    }

    public function getAge(): int
    {
        return $this->age;
    }

    public function setCost(float $cost): self
    {
        return $this->cost = $cost;
        return $this;
    }

    public function getCost(): float
    {
        return $this->cost;
    }

}
```