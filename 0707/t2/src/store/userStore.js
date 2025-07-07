import { create } from 'zustand';

const useUserStore = create((set, get) => ({
  user: {},
  fetchUser: async () => {
    const fakeApi = () =>
      new Promise(resolve =>
        setTimeout(() =>
          resolve({ firstName: '张', lastName: '三', age: 28 }), 1000)
      );
    const userData = await fakeApi();
    set({ user: userData });
  },
  get isUserLoaded() {
    return Object.keys(get().user).length > 0;
  },
  get userFullName() {
    const { firstName = '', lastName = '' } = get().user;
    return `${firstName} ${lastName}`.trim();
  }
}));

useUserStore.subscribe(
  state => state.user,
  (user) => {
    console.log('用户信息变化:', user);
  }
);

export default useUserStore;
