function ProductCard({ product, onRemove, onUpdateQty }) {
  return (
    <div style={{ 
      display: 'flex', 
      justifyContent: 'space-between', 
      padding: '10px', 
      borderBottom: '1px solid #ddd',
      alignItems: 'center' 
    }}>
      <div>
        <strong>{product.name}</strong> — {product.price} руб.
      </div>
      
      <div style={{ display: 'flex', alignItems: 'center', gap: '10px' }}>
        <button onClick={() => onUpdateQty(product.id, -1)}>−</button>
        <span>{product.qty}</span>
        <button onClick={() => onUpdateQty(product.id, 1)}>+</button>
        <button onClick={() => onRemove(product.id)} style={{ color: 'red' }}>🗑️</button>
      </div>
    </div>
  );
}

export default ProductCard;