import Logo from './Logo';
import Nav from './Nav';
import ToggleThemeButton from './ToggleThemeButton';

export default function Header({ theme, setTheme }) {
  return (
    <header style={{ display: 'flex', justifyContent: 'space-between', padding: '20px', borderBottom: '1px solid #ccc' }}>
      <Logo />
      <Nav />
      <div style={{ display: 'flex', alignItems: 'center', gap: '10px' }}>
        <div className="small-logo">🔹</div>
        <ToggleThemeButton theme={theme} setTheme={setTheme} />
      </div>
    </header>
  );
}