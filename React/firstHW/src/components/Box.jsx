function Box({ children }) {
  return (
    <div style={{ border: '2px dotted blue', padding: '20px', borderRadius: '8px' }}>
      {children}
    </div>
  );
}
export default Box;