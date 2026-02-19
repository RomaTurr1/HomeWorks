export default function ProductItem({product,index}){
    return(
        <div style={{border: "1px solid black",
                    borderRadius: 10,
                    padding: 12,
                    marginBottom: 10,
        }}>
            <div>
                <b>{index + 1}.</b>
                {product.title}
                <div>
                    Цена: {product.price}
                </div>
            </div>
        </div>
    )
}