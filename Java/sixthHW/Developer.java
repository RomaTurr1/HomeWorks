package Java.sixthHW;

public class Developer extends Employee {
    private String language;

    public Developer(String name, double salary, String language) {
        super(name, salary);
        setLanguage(language);
    }
    @Override
    public void info() {
        super.info();
        System.out.println("Language: " + language);
    }
    @Override
    public void work() {
        System.out.println(getName() + " пишет код на " + language);
    }

    public String getLanguage() {
    return language;
    }

    public void setLanguage(String language) {
    if (language == null || language.isEmpty()) {
        System.out.println("Ошибка: язык не может быть пустым");
        return;
    }
    this.language = language;
    }
}
