import ProductItem from "./ProductItem"


export default function ProductList({items}){
    return(
        <>  
            {items.map((product, index) =>(
                <ProductItem 
                key = {product.id}
                product = {product}
                index = {index}
                />
            ))}
        </>
    )
}