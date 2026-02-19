import Header from "./components/Header"
import Main from "./components/Main"
import ProductList from "./components/ProductList"
import { useState } from "react"
import Todo from "./components/Todo"

const products = [
    {id: 1, title: "Keyboard", price:12000},
    {id: 2, title: "Mouse", price:7000},
    {id: 3, title: "Headphones", price:35000}
]
// let count = 0

function App(){
    const [count, setCount] = useState(0)
    const [show, setShow] = useState(true)
    const [name, setName] = useState("")
    const [names, setNames] = useState(["Alex", "Dana"])
    const [value, setValue] = useState("")

    const addName = () =>{
        setNames([...names, value])
        setValue("")
    }


    return(
        <>

            {/* <div>
                    <p>{count}</p>
                    <button onClick={() => count++}>Btn</button>
                </div> */}


                <div>
                    <p>{count}</p>
                    <button onClick={() => setCount(count + 1)}>+</button>
                    <button onClick={() => setCount(count - 1)}>-</button>
                </div>

                <div>
                    <button onClick={() => setShow(!show)}>Toggle</button>
                    {show && <p>Hello</p>}
                </div>

                <div>
                    <input
                    value={name}
                    onChange={(e) => setName(e.target.value)}
                    />
                    <p>Hello, {name}</p>
                </div>

                <div>
                    <input 
                    type="text"
                    value={value}
                    onChange={(e) => {setValue(e.target.value)}}
                    placeholder="Compose some name..."
                     />
                     <button onClick={addName}>Add</button>
                     {names.map((name,index) =>(
                        <p key={index}>{name}</p>
                     ))}
                </div>

            <Header></Header>
            <Main>
                
            </Main>
            <div style={{padding:20}}>
                <h1>Товары</h1>
                <ProductList items={products}/>
            </div>
            <Todo/>
            
        </>
    )
}
export default App