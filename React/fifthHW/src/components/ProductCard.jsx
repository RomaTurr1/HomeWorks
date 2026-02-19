export default function ProductCard({ product }) {
  return (
    <div className="product-card" style={{ 
      background: 'var(--card-bg)', 
      padding: '15px', 
      borderRadius: '8px',
      border: '1px solid #ddd'
    }}>
      <img src={product.image} alt={product.title} style={{ width: '100%', height: '150px', objectFit: 'contain' }} />
      <h4 style={{ fontSize: '14px', height: '40px', overflow: 'hidden' }}>{product.title}</h4>
      <p style={{ fontSize: '12px' }}>{product.description.slice(0, 60)}...</p>
      <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
        <span style={{ fontWeight: 'bold' }}>${product.price}</span>
        <span>⭐ {product.rating?.rate || 'No rating'}</span>
      </div>
      <button onClick={() => alert(product.title)} style={{ marginTop: '10px', width: '100%' }}>MORE</button>
    </div>
  );
}