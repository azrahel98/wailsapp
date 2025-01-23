<template>
  <div class="card">
    <div class="card-header" v-if="title">
      <h3 class="card-title">{{ title }}</h3>
    </div>
    <div class="card-body border-bottom py-2">
      <div class="d-flex">
        <div class="ms-auto">
          <div class="input-icon">
            <span class="input-icon-addon">
              <IconSearch class="icon" />
            </span>
            <input
              type="text"
              v-model="searchQuery"
              class="form-control form-control-sm form-control-rounded"
              placeholder="Busca..."
            />
          </div>
        </div>
      </div>
    </div>
    <div class="table-responsive">
      <table class="table table-sm card-table">
        <thead>
          <tr>
            <th
              v-for="column in columns"
              :key="column.field"
              @click="sortBy(column.field)"
              class="text-nowrap cursor-pointer text-start"
              :class="{ 'text-primary': sortColumn === column.field }"
            >
              {{ column.title }}
              <span v-if="sortColumn === column.field" class="ms-1">
                {{ sortOrder === 'asc' ? '↑' : '↓' }}
              </span>
            </th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="row in paginatedData" :key="row.id" class="fs-5" @click="ruta(row.Nombre)">
            <td v-for="column in columns" :key="column.field" class="text-start fw-medium">
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
import { IconSearch } from '@tabler/icons-vue'
import { computed, ref, watch } from 'vue'

interface Column {
  field: string
  title: string
}

interface Row {
  [key: string]: any
}

const props = defineProps<{
  rows: Row[]
  columns: Column[]
  title?: string
  initialPageSize?: number
}>()

const currentPage = ref(1)
const sortColumn = ref('')
const sortOrder = ref<'asc' | 'desc'>('asc')
const searchQuery = ref('')
const pageSize = ref(props.initialPageSize || 10)

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
  const end = start + pageSize.value
  return sortedData.slice(start, end)
})

const sortBy = (column: string) => {
  if (sortColumn.value === column) {
    sortOrder.value = sortOrder.value === 'asc' ? 'desc' : 'asc'
  } else {
    sortColumn.value = column
    sortOrder.value = 'asc'
  }
}

const ruta = async (area: string) => {
  await router.push({
    name: 'inofarea',
    params: {
      area: area
    }
  })
}

watch([searchQuery, pageSize], () => {
  currentPage.value = 1
})
</script>

<style lang="scss" scoped>
table {
  th {
    font-weight: 500;
    text-transform: none;
    font-size: 0.76rem;
    padding: 0.5rem;
    border-bottom-width: 1px;
    background: #f8fafc;
  }

  td {
    padding: 0.5rem;
    font-size: 0.7rem;
    border-bottom: 1px solid #e2e8f0;
  }
  tr:hover {
    background-color: var(--tblr-primary-fg);
    font-weight: 600;
  }
}
</style>
