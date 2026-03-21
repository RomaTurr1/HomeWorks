package Java.sixthHW;

public class Employee {
    private String name;
    private double salary;

    public Employee(String name, double salary){
        this.name = name;
        this.salary = salary;
    }
    public void info(){
        System.out.println("Сотрудник:" + name);
        System.out.println("Зарплата:" + salary);
    }
    public void work(){
        System.out.println(name + "Выполняет работу");
    }

    public String getName() {
    return name;
    }

    public double getSalary() {
    return salary;
    }
    public void setName(String name) {
    if (name == null || name.isEmpty()) {
        System.out.println("Ошибка: имя не может быть пустым");
        return;
    }
    this.name = name;
    }

    public void setSalary(double salary) {
    if (salary < 0) {
        System.out.println("Ошибка: зарплата не может быть отрицательной");
        return;
    }
    this.salary = salary;
    }

    public void raiseSalary(double amount) {
    if (amount <= 0) {
        System.out.println("Ошибка: некорректное повышение");
        return;
    }
    this.salary += amount;
    }
}


