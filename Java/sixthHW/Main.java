package Java.sixthHW;


public class Main {
        //   public static void main(String[] args){
        //     // Car car1 = new Car();
        //     // car1.brand = "Toyota";
        //     // car1.color = "White";
        //     // car1.speed = 20;

        //     // Car car2 = new Car();
        //     // car2.brand = "BMW";
        //     // car2.color = "Black";
        //     // car2.speed = 50;

        //     // car1.info();
        //     // car1.accelerate(50);
        //     // car1.brake(20);
        //     // System.out.println("---------");
        //     // car2.info();
        //     // car2.accelerate(50);
        //     // car2.brake(20);

        //     Car[] cars = new Car[3];
        //     cars[0] = new Car();
        //     cars[0].brand = "Toyota";
        //     cars[0].color = "White";
        //     cars[0].speed = 20;

        //     cars[1] = new Car();
        //     cars[1].brand = "BMW";
        //     cars[1].color = "Black";
        //     cars[1].speed = 50;

        //     cars[2] = new Car();
        //     cars[2].brand = "Honda";
        //     cars[2].color = "Blue";
        //     cars[2].speed = 35;

        //     for (Car c : cars){
        //         c.info();

        //         System.out.println("----------------");
        //     }

        //         BankAccount acc = new BankAccount();
        //         acc.setOwner("Amina");
        //         acc.setBalance(5000);

        //         System.out.println(acc.getOwner());
        //         System.out.println(acc.getBalance());

        //         acc.setBalance(-7000);

        //         System.out.println(acc.getBalance());


        //         Student s1 = new Student("Ali", 18, 92.5);
        //         Student s2 = new Student("Diana", 19, 88.0);

        //         s1.info();
        //         s2.info();


        //         Developer d1 = new Developer("Alex", 5000, "Python");
        //         Designer de1 = new Designer("Sasha", 3500, "Figma");
        //         Manager m1 = new Manager("Arman", 7000, 15);

        //         d1.info();
        //         d1.work();

        //         de1.info();
        //         de1.work();

        //         m1.info();
        //         m1.work();

// HomeWork :
    public static void main(String[] args) {

    Developer d = new Developer("Alex", 80000, "Java");
    Designer des = new Designer("Anna", 60000, "Figma");
    Manager m = new Manager("Oleg", 100000, 5);

    d.info();
    d.work();

    des.info();
    des.work();

    m.info();
    m.work();

    System.out.println("\n--- Проверка setter ---");

    d.setLanguage("");
    des.setTool("");
    m.setTeamSize(-1);

    d.setName("");
    d.setSalary(-500);

    System.out.println("\n--- Повышение зарплаты ---");

    d.raiseSalary(5000);
    d.info();
}
}
