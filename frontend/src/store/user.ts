import { defineStore } from "pinia"

export const useCounterStore = defineStore('counter', {
    state: () => ({ admin: false, name: '' }),
    getters: {
      isAdmin: (state) => state.admin ,
    },
    actions: {
      create() {
        
      },
    },
  })