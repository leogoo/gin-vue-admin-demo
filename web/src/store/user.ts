import { defineStore } from 'pinia'

export const userStore = defineStore('user', {
  state: () => ({ count: 1, name: 'bob'}),
  getters: {
    doubleCount: (state) => state.count * 2
  },
  actions: {
    increment() {
      this.count++
    }
  }
})