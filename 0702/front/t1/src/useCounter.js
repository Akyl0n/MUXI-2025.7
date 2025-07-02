import { useState, useEffect, useCallback } from 'react';

function useCounter(initialValue = 0) {
  const [count, setCount] = useState(initialValue);

  const increment = useCallback(() => setCount(c => c + 1), []);
  const decrement = useCallback(() => setCount(c => c - 1), []);
  const reset = useCallback(() => setCount(initialValue), [initialValue]);

  useEffect(() => {
    console.log('计数值变化:', count);
  }, [count]);

  return { count, increment, decrement, reset };
}

export default useCounter;
