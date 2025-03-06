<template>
  <div class="card">
    <div class="card-header">
      <h3 class="card-title">Organos / Unidades</h3>
    </div>
    <div class="card-body border-bottom py-1">
      <div class="d-flex">
        <div class="ms-auto text-secondary">
          Buscar:
          <div class="ms-2 d-inline-block">
            <input type="text" v-model="searchQuery" class="form-control" />
          </div>
        </div>
      </div>
    </div>
    <div class="table-responsive">
      <table class="table card-table table-vcenter text-nowrap datatable">
        <thead>
          <tr>
            <th
              v-for="column in columns"
              :key="column.field"
              @click="sortBy(column.field)"
              class="cursor-pointer"
            >
              {{ column.title }}
              <span v-if="sortColumn === column.field">
                {{ sortOrder === 'asc' ? '↑' : '↓' }}
              </span>
            </th>
          </tr>
        </thead>
        <tbody>
          <tr
            v-for="row in paginatedData"
            :key="row.id"
            class="clickable-row"
            @click="gerArea(row.Nombre)"
          >
            <td v-for="column in columns" :key="column.field">
              {{ row[column.field] }}
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script setup lang="ts">
import { router } from '@router/router'
import { computed, ref } from 'vue'

interface Column {
  field: string
  title: string
}
interface Row {
  [key: string]: any
}

const props = defineProps<{ rows: Row[]; columns: Column[]; initialPageSize?: number }>()

const currentPage = ref(1)
const sortColumn = ref('')
const sortOrder = ref<'asc' | 'desc'>('asc')
const searchQuery = ref('')
const pageSize = ref(props.initialPageSize || 20)

const gerArea = async (area: string) => {
  await router.push({
    name: 'inofarea',
    params: {
      area: area
    }
  })
}

const filteredRows = computed(() => {
  return props.rows.filter((row) =>
    Object.values(row).some((value) =>
      String(value).toLowerCase().includes(searchQuery.value.toLowerCase())
    )
  )
})

const paginatedData = computed(() => {
  let sortedData = [...filteredRows.value]
  if (sortColumn.value) {
    sortedData.sort((a, b) => {
      if (a[sortColumn.value] < b[sortColumn.value]) return sortOrder.value === 'asc' ? -1 : 1
      if (a[sortColumn.value] > b[sortColumn.value]) return sortOrder.value === 'asc' ? 1 : -1
      return 0
    })
  }
  const start = (currentPage.value - 1) * pageSize.value
  return sortedData.slice(start, start + pageSize.value)
})

const sortBy = (column: string) => {
  if (sortColumn.value === column) {
    sortOrder.value = sortOrder.value === 'asc' ? 'desc' : 'asc'
  } else {
    sortColumn.value = column
    sortOrder.value = 'asc'
  }
}
</script>

<style scoped>
.table-responsive {
  max-height: 30vh;
  height: 100%;
  overflow: scroll;
}
.table th {
  cursor: pointer;
}
.clickable-row:hover {
  background-color: rgba(0, 0, 0, 0.05);
  cursor: pointer;
}
</style>
