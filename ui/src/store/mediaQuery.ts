import { ref } from 'vue'
import { defineStore } from 'pinia'
import { mediaQuery } from '@/utils'

export const useMediaQuery = defineStore('mediaQuery', () => {
  const sm = ref(false)
  const md = ref(false)
  const lg = ref(false)
  const xl = ref(false)
  const doubleXL = ref(false)

  mediaQuery('(min-width: 640px)', (mq) => {
    sm.value = mq.matches
  })
  mediaQuery('(min-width: 768px)', (mq) => {
    md.value = mq.matches
  })
  mediaQuery('(min-width: 1024px)', (mq) => {
    lg.value = mq.matches
  })
  mediaQuery('(min-width: 1280px)', (mq) => {
    xl.value = mq.matches
  })
  mediaQuery('(min-width: 1536px)', (mq) => {
    doubleXL.value = mq.matches
  })

  return { sm, md, lg, xl, doubleXL }
})
