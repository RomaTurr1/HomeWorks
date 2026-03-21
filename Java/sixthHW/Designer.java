package Java.sixthHW;

public class Designer extends Employee{
    private String tool;

    public Designer(String name, double salary, String tool) {
        super(name, salary);
        setTool(tool);
    }   
    @Override
    public void info() {
        super.info();
        System.out.println("Tool: " + tool);
    }
    @Override
    public void work() {
        System.out.println(getName() + " делает дизайн в " + tool);
    }

    public String getTool() {
    return tool;
    }

    public void setTool(String tool) {
    if (tool == null || tool.isEmpty()) {
        System.out.println("Ошибка: инструмент не может быть пустым");
        return;
    }
    this.tool = tool;
    }
}
