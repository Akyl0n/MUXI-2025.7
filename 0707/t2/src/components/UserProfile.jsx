import React from 'react';
import useUserStore from '../store/userStore';

const UserProfile = () => {
  const user = useUserStore(state => state.user);
  const fetchUser = useUserStore(state => state.fetchUser);
  const isUserLoaded = useUserStore(state => state.isUserLoaded);
  const userFullName = useUserStore(state => state.userFullName);

  return (
    <div style={{ padding: 24, border: '1px solid #eee', borderRadius: 8, maxWidth: 320 }}>
      <h2>用户信息</h2>
      <button onClick={fetchUser} disabled={isUserLoaded}>
        {isUserLoaded ? '用户已加载' : '加载用户信息'}
      </button>
      <div style={{ marginTop: 16 }}>
        {isUserLoaded ? (
          <>
            <div>姓名：{userFullName}</div>
            <div>年龄：{user.age}</div>
          </>
        ) : (
          <div>Loading user...</div>
        )}
      </div>
    </div>
  );
};

export default UserProfile;
