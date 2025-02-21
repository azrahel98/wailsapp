<template>
  <div class="card">
    <div class="card-header">
      <h3 class="card-title">Cumpleaños</h3>
    </div>
    <div class="table-responsive table-responsive-sm">
      <table class="table card-table table-vcenter table-hover">
        <tbody>
          <tr v-for="x in filteredInfo(info)" @click="getperfil(x.Dni)">
            <td class="w-100">
              <a @click="getperfil(x.Dni)" class="small text-reset">{{ x.Nombre }}</a>
            </td>
            <td class="text-nowrap text-secondary fw-medium" v-if="istoday(x.Nacimiento)">
              <IconCalendar class="icon text-primary" />
              {{ format(addDays(parseISO(x.Nacimiento), 1), 'MMMM dd, yyyy') }}
            </td>
            <td class="text-nowrap text-secondary" v-else>
              <IconCalendar class="icon icon-1" />
              {{ format(addDays(parseISO(x.Nacimiento), 1), 'MMMM dd, yyyy') }}
            </td>
            <td class="text-nowrap">
              <a class="text-secondary">
                <IconCake class="icon icon-1" />
                {{ x.Edad }}
              </a>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { router } from '@router/router'
import { IconCake, IconCalendar, IconCheck, IconCooker } from '@tabler/icons-vue'
import { addDays, format, parseISO, startOfDay } from 'date-fns'

defineProps({
  info: { type: Array<any>, required: true }
})

const istoday = (x: string): boolean => {
  const today = startOfDay(new Date())
  const birthday = addDays(parseISO(x), 1)
  return today.getDate() == birthday.getDate()
}

const filteredInfo = (data: Array<any>) => {
  const today = startOfDay(new Date())
  return data.filter((x) => {
    const birthday = addDays(parseISO(x.Nacimiento), 1)
    return today.getDate() <= birthday.getDate()
  })
}

const getperfil = async (dni: string) => {
  await router.push({
    name: 'perfil',
    params: {
      dni: dni.toString()
    }
  })
}
</script>

<style lang="scss" scoped>
.table-responsive {
  max-height: 20vh;
  overflow: scroll;
}
</style>
