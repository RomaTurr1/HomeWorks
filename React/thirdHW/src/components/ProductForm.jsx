import { useState } from 'react';

function ProductForm({ onAdd }) {
  const [form, setForm] = useState({ name: "", price: "", qty: 1 });

  const handleSubmit = (e) => {
    e.preventDefault();
    if (!form.name || form.price <= 0) return;
    
    // Передаем данные наверх
    onAdd({ ...form, id: Date.now(), price: Number(form.price), qty: Number(form.qty) });
    
    // Очистка
    setForm({ name: "", price: "", qty: 1 });
  };

  return (
    <form onSubmit={handleSubmit} style={{ display: 'flex', gap: '10px', marginBottom: '20px' }}>
      <input 
        placeholder="Название" 
        value={form.name} 
        onChange={e => setForm({...form, name: e.target.value})} 
      />
      <input 
        type="number" 
        placeholder="Цена" 
        value={form.price} 
        onChange={e => setForm({...form, price: e.target.value})} 
      />
      <button type="submit">Добавить</button>
    </form>
  );
}

export default ProductForm;