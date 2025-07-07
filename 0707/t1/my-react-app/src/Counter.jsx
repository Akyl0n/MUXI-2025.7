import { useCounterStore } from './store'

export function Counter() {
  const { count, increment, decrement, isPositive } = useCounterStore()
  
  return (
    <div style={{ 
      fontFamily: 'sans-serif',
      maxWidth: '400px',
      margin: '0 auto',
      padding: '20px',
      textAlign: 'center'
    }}>
      <h1>Counter App</h1>
      <div style={{ 
        fontSize: '2rem',
        margin: '20px 0'
      }}>
        Count: {count}
      </div>
      <div style={{ 
        margin: '20px 0',
        color: isPositive() ? 'green' : 'red'
      }}>
        {isPositive() ? 'Count is positive' : 'Count is not positive'}
      </div>
      <div style={{ 
        display: 'flex',
        justifyContent: 'center',
        gap: '10px'
      }}>
        <button 
          onClick={decrement}
          style={{
            padding: '10px 20px',
            fontSize: '1rem',
            cursor: 'pointer'
          }}
        >
          Decrement
        </button>
        <button 
          onClick={increment}
          style={{
            padding: '10px 20px',
            fontSize: '1rem',
            cursor: 'pointer'
          }}
        >
          Increment
        </button>
      </div>
    </div>
  )
}