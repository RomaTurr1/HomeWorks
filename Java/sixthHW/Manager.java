package Java.sixthHW;

public class Manager extends Employee{
    private int teamSize;

    public Manager(String name, double salary, int teamSize) {
        super(name, salary);
        setTeamSize(teamSize);
    }
    @Override
    public void info() {
        super.info();
       System.out.println("Team size: " + teamSize);
    }
    @Override
    public void work() {
        System.out.println(getName() + " управляет командой из " + teamSize + " человек");
    }

    public int getTeamSize() {
    return teamSize;
    }

    public void setTeamSize(int teamSize) {
    if (teamSize < 0) {
        System.out.println("Ошибка: размер команды не может быть отрицательным");
        return;
    }
    this.teamSize = teamSize;
    }
}

