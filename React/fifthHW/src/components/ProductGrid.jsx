import ProductCard from './ProductCard';

export default function ProductGrid({ products }) {
  return (
    <div style={{ 
      display: 'grid', 
      gridTemplateColumns: 'repeat(4, 1fr)', 
      gap: '20px', 
      marginTop: '20px' 
    }}>
      {products.map(item => <ProductCard key={item.id} product={item} />)}
    </div>
  );
}