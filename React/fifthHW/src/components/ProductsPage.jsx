import { useState, useEffect } from 'react';
import SearchBar from './SearchBar';
import ProductGrid from './ProductGrid';

export default function ProductsPage() {
  const [products, setProducts] = useState([]); // Все 8 товаров
  const [filteredProducts, setFilteredProducts] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  // useEffect №1: Загрузка товаров
  useEffect(() => {
    fetch('https://fakestoreapi.com/products')
      .then(res => {
        if (!res.ok) throw new Error('Ошибка загрузки данных');
        return res.json();
      })
      .then(data => {
        const slicedData = data.slice(0, 8); // Ровно 8 карточек
        setProducts(slicedData);
        setFilteredProducts(slicedData);
        setLoading(false);
      })
      .catch(err => {
        setError(err.message);
        setLoading(false);
      });
  }, []);

  const handleSearch = (searchTerm) => {
    const filtered = products.filter(p => 
      p.title.toLowerCase().includes(searchTerm.toLowerCase())
    );
    setFilteredProducts(filtered);
  };

  if (loading) return <h2>Загрузка...</h2>;
  if (error) return <h2 style={{ color: 'red' }}>{error}</h2>;

  return (
    <main style={{ padding: '20px' }}>
      <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
        <h2>Products</h2>
        <SearchBar onSearch={handleSearch} />
      </div>
      <ProductGrid products={filteredProducts} />
    </main>
  );
}