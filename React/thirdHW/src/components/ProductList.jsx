import ProductCard from './ProductCard';

function ProductList({ products, onRemove, onUpdateQty }) {
  return (
    <div>
      {products.map(product => (
        <ProductCard 
          key={product.id} 
          product={product} 
          onRemove={onRemove} 
          onUpdateQty={onUpdateQty} 
        />
      ))}
    </div>
  );
}

export default ProductList;