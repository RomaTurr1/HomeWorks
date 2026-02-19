function ProductCard({ title, price, inStock }) {
  return (
    <div style={{ color: inStock ? 'black' : 'gray' }}>
      <h4>{title}</h4>
      <p>Цена: ${price}</p>
      <p>{inStock ? "✅ В наличии" : "❌ Нет на складе"}</p>
    </div>
  );
}
export default ProductCard;