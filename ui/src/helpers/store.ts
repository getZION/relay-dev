import create from 'zustand'

const useStore = create(() => {
  return {
    router: null,
    dom: null,
    pubkey: '',
    balance: 0,
  }
})

export default useStore
