import { createSlice } from '@reduxjs/toolkit'

export interface ProfileSideBarState {
  open: boolean
}

const initialState: ProfileSideBarState = {
  open: false,
}

export const profileSideBarSlice = createSlice({
  name: 'ProfileSideBar',
  initialState,
  reducers: {
    open: (state) => {
      document.body.style.overflow = "hidden";
      state.open = true
    },
    close: (state) => {
      document.body.style.overflow = 'unset';
      state.open = false
    },
  },
})

// Action creators are generated for each case reducer function
export const { open, close } = profileSideBarSlice.actions

export default profileSideBarSlice.reducer