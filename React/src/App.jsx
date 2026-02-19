import Header from './components/Header';
import Main from './components/Main';
import Footer from './components/Footer';
import UserCard from './components/UserCard';
import ProductCard from './components/ProductCard';
import Box from './components/Box';

function App() {
  const currentYear = new Date().getFullYear();
  const userName = "Твоё Имя";
  const isLoggedIn = true; // Для примера тернарного оператора

  return (
    <div>
      <Header />
      
      <Main>
        <h1>Мой первый React-проект</h1>
        <p>Автор: {userName}</p>
        <p>Год: {currentYear}</p>
        
        {/* Условный рендеринг */}
        <p>{isLoggedIn ? "Добро пожаловать обратно!" : "Пожалуйста, войдите"}</p>

        <hr />

        <h2>Список пользователей</h2>
        <UserCard name="Алексей" age={25} city="Москва" isStudent={true} />
        <UserCard name="Мария" age={30} city="Алматы" isStudent={false} />
        <UserCard name="Джон" age={19} city="Нью-Йорк" isStudent={true} />

        <hr />

        <h2>Наши товары</h2>
        <ProductCard title="Ноутбук" price={1000} inStock={true} />
        <ProductCard title="Смартфон" price={500} inStock={false} />
        <ProductCard title="Наушники" price={150} inStock={true} />

        <hr />

        <Box>
          <h3>Это контент внутри Box</h3>
          <p>Я передан через props.children!</p>
        </Box>
      </Main>

      <Footer />
    </div>
  );
}

export default App;