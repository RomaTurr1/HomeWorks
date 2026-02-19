import Card from "./Card"
import Text from "./Text"
export default function Main(){
    const names = ["Aliya", "Magzhan", "Diana"]

    return(
        <>
        <Card title = "Iphone 16" desc = "Created by Apple company"/>
        <Card title = "Samsung S24" desc = "Created by Samsung company"/>
        <Card title = "Pixel 10" desc = "Created by Google company"/>
        
        <div>
          {names.map((name,index) =>(
            <Text key = {index} text={name}/>
          ))}
        </div>
        </>
    )
}