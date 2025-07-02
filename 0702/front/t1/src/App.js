import useCounter from './useCounter';
import './App.css';

function App() {
  const { count, increment, decrement, reset } = useCounter(0);
  return (
    <div className="counter-container">
      <div className="counter-value">当前计数： {count}</div>
      <div className="button-group">
        <button className="counter-btn" onClick={increment}>➕ 增加</button>
        <button className="counter-btn" onClick={decrement}>➖ 减少</button>
        <button className="counter-btn" onClick={reset}>↺ 重置</button>
      </div>
    </div>
  );
}

export default App;
