import profileSideBarReducer from '../user/components/ProfileSideBar/ProfileSideBarSlice';
import { configureStore } from '@reduxjs/toolkit'

const store = configureStore({
    reducer: {
        profileSideBar: profileSideBarReducer
    },
})

// Infer the `RootState` and `AppDispatch` types from the store itself
export type RootState = ReturnType<typeof store.getState>
// Inferred type: {posts: PostsState, comments: CommentsState, users: UsersState}
export type AppDispatch = typeof store.dispatch

export default store;