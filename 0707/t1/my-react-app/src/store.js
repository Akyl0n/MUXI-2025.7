import { create } from 'zustand'
import { persist, createJSONStorage } from 'zustand/middleware'

export const useCounterStore = create(
  persist(
    (set, get) => ({
      count: 0,
      increment: () => set((state) => ({ count: state.count + 1 })),
      decrement: () => set((state) => ({ count: state.count - 1 })),
      isPositive: () => get().count > 0,
    }),
    {
      name: 'counter-storage', 
      storage: createJSONStorage(() => localStorage),
    }
  )
)