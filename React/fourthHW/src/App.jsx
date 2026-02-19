import { useState } from 'react';
import styles from './App.module.css';
import ProductForm from './components/ProductForm';
import ProductList from './components/ProductList';

function App() {
  const [products, setProducts] = useState([]);
  const [filter, setFilter] = useState('all'); // 'all', 'bought', 'active'

  // Добавление товара (не забываем про bought: false)
  const addProduct = (newProduct) => {
    setProducts(prev => [{ ...newProduct, bought: false }, ...prev]);
  };

  // ✅ ГЛАВНОЕ: Toggle через map без мутаций
  const toggleBought = (id) => {
    setProducts(prev => 
      prev.map(p => p.id === id ? { ...p, bought: !p.bought } : p)
    );
  };

  const removeProduct = (id) => {
    setProducts(prev => prev.filter(p => p.id !== id));
  };

  // Логика фильтрации (Доп. задание ⭐)
  const filteredProducts = products.filter(p => {
    if (filter === 'bought') return p.bought;
    if (filter === 'active') return !p.bought;
    return true;
  });

  return (
    <div className={styles.container}>
      <h2>🛒 Mini Shop Cart</h2>
      
      <ProductForm onAdd={addProduct} />

      {/* Фильтры ⭐ */}
      <div className={styles.filters}>
        <button onClick={() => setFilter('all')} className={filter === 'all' ? styles.active : ''}>Все</button>
        <button onClick={() => setFilter('active')} className={filter === 'active' ? styles.active : ''}>Активные</button>
        <button onClick={() => setFilter('bought')} className={filter === 'bought' ? styles.active : ''}>Купленные</button>
      </div>
      
      <ProductList 
        products={filteredProducts} 
        onRemove={removeProduct} 
        onToggle={toggleBought} 
      />
    </div>
  );
}

export default App;