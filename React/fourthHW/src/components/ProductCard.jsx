function ProductCard({ product, onRemove, onToggle }) {
  // Определяем стили в зависимости от состояния bought
  const itemStyle = {
    textDecoration: product.bought ? 'line-through' : 'none',
    color: product.bought ? '#9ca3af' : '#2c3e50',
    opacity: product.bought ? 0.7 : 1,
    transition: '0.3s'
  };

  return (
    <div style={{ 
      display: 'flex', 
      justifyContent: 'space-between', 
      padding: '12px', 
      borderBottom: '1px solid #eee',
      alignItems: 'center',
      backgroundColor: product.bought ? '#f9fafb' : '#fff'
    }}>
      <div style={{ display: 'flex', alignItems: 'center', gap: '12px' }}>
        {/* Чекбокс для переключения */}
        <input 
          type="checkbox" 
          checked={product.bought} 
          onChange={() => onToggle(product.id)} 
        />
        
        <div style={itemStyle}>
          <strong>{product.name}</strong> 
          {product.bought && <span style={{ marginLeft: '8px', fontSize: '12px', color: 'green' }}>✔ Куплено</span>}
        </div>
      </div>
      
      <button 
        onClick={() => onRemove(product.id)} 
        style={{ background: 'none', border: 'none', cursor: 'pointer', color: '#ef4444' }}
      >
        🗑️
      </button>
    </div>
  );
}

export default ProductCard;