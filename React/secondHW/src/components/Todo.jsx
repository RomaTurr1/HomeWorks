import { useState } from "react";

function Todo(){
const [task, setTask] = useState("")
const [tasks, setTasks] = useState([])

const [editIndex, setEditIndex] = useState(null)
const [editValue, setEditValue] = useState("")

function addTask(){
    if(task === "") return

    setTasks([...tasks, task])
    setTask("")
}

function removeTask(index){
    setTasks(tasks.filter((_, i) =>{
        i !== index
    }))
}
function startEdit(index){
    setEditIndex(index)
    setEditValue(tasks[index])
}
function saveEdit(){
    const newTasks = [...tasks]
    newTasks[editIndex] = editValue
    
    setTasks(newTasks)
    setEditIndex(null)
    setEditValue("")
}

return(
    <div>
        <h2>To-Do list</h2>

        <input
         type="text" 
         placeholder="Введите заметку"
         value={task}
         onChange={(e) => setTask(e.target.value)}
         />
         <button onClick={addTask}>Добавить</button>

         <ul>
            {tasks.map((item, index) =>(
                <li key={index}>
                    {editIndex === index ? (
                        <>
                            <input 
                            type="text"
                             value={editValue}
                             onChange={(e) => setEditValue(e.target.value)}
                             />
                             <button onClick={saveEdit}>Сохранить</button>
                        </>

                    ): (
                        <>
                            {item}
                            <button onClick={() => startEdit(index)}>✏</button>
                            <button onClick={() => removeTask(index)}>🗑</button>
                        </>
                    )}
                </li>
            ))}
         </ul>
    </div>
)


}
export default Todo