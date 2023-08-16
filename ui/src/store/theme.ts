import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useThemeStore = defineStore('theme', () => {
  const theme = ref({
    headerBgColor: '#22232e',
    sidebarBgColor: '#1e1f29',
    mainBgColor: '#191a23',
    textColor: '#c2cad1',
    textBtnHoverBgColor: '#e6e7e90a',
    btnTextPrimaryColor: '#191a23',
    primaryColor: '#8276c9',
    btnHoverBgColor: '#756ab5',
    sideBarInputBgColor: '#252733',
    treeNodeHoverBgColor: '#8276c914',
    splitBorderColor:'#252431'
  })
  return { theme }
})
