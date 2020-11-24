export const state = () => ({
  admin: false
})

export const getters = {
  adminStatus (state) {
    return state.admin
  }
}

export const mutations = {
  adminTrue(state) {
    state.admin = true
  },
  adminFalse(state) {
    state.admin = false
  }
}
