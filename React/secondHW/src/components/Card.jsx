export default function Card({title,desc}){
    return(
        <>
            <div className="card">
                <h1>{title}</h1>
                <p>{desc}</p>
                <button>Buy</button>
            </div>
        </>
    )
}