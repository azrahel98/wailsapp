//@ts-nocheck
import jsPDF from 'jspdf'
import 'jspdf-autotable'

export const makepdf = () => {
  const doc = new jsPDF('portrait', 'pt', 'a4')

  // Márgenes y configuraciones iniciales
  const pageWidth = doc.internal.pageSize.getWidth()
  const pageHeight = doc.internal.pageSize.getHeight() // Definir la altura de la página
  const margin = 10

  // Encabezado
  doc.setFontSize(12)
  doc.addImage('../assets/vue.svg', 'JPEG', margin, margin, 60, 60) // Logo
  doc.setFontSize(10)
  doc.text('Municipalidad de Villa el Salvador', margin, margin+20)
  doc.setFontSize(16)
  doc.text('BOLETA DE PAGO - PLANILLA', pageWidth / 2, margin + 40, { align: 'center' })
  doc.setFontSize(10)
  doc.text('Nº Documento: 00001', pageWidth - margin, margin + 20, { align: 'right' })
  doc.text('Fecha: 2024-01-31', pageWidth - margin, margin + 40, { align: 'right' })
  doc.text('Período: Enero 2024', pageWidth / 2, margin + 60, { align: 'center' })

  // Información del empleador y empleado (dos columnas)
  const sectionStartY = margin + 90
  const columnWidth = (pageWidth - 2 * margin) / 2
  doc.setFontSize(10)

  // Columna 1
  doc.text('Entidad: XXXXXXXXXX', margin, sectionStartY)
  doc.text('Empleador: XXXXXXX', margin, sectionStartY + 15)
  doc.text('RUC: 12345678901', margin, sectionStartY + 30)

  // Columna 2
  doc.text('Rubro de Financiamiento: XXXXXXX', margin + columnWidth, sectionStartY)
  doc.text('Meta Presupuestal: 10101', margin + columnWidth, sectionStartY + 15)
  doc.text('Unidad Orgánica: Dirección General', margin + columnWidth, sectionStartY + 30)

  // Información personal del empleado (dos columnas)
  const employeeStartY = sectionStartY + 60
  doc.text('Doc. Identidad: 12345678', margin, employeeStartY)
  doc.text('Apellidos y Nombres: Juan Perez', margin, employeeStartY + 15)
  doc.text('Fecha de Ingreso: 2020-01-15', margin, employeeStartY + 30)
  doc.text('Régimen de Pensión: AFP', margin, employeeStartY + 45)
  doc.text('Administradora de Pensión: XXX', margin, employeeStartY + 60)
  doc.text('CUSPP: XXXXXXXXXX', margin, employeeStartY + 75)

  doc.text('Código AIRHSP: 000000', margin + columnWidth, employeeStartY)
  doc.text('Establecimiento: Central', margin + columnWidth, employeeStartY + 15)
  doc.text('Régimen Laboral: CAS', margin + columnWidth, employeeStartY + 30)
  doc.text('Condición: Permanente', margin + columnWidth, employeeStartY + 45)
  doc.text('Grupo Ocupacional: Administrativo', margin + columnWidth, employeeStartY + 60)
  doc.text('Cargo: Analista', margin + columnWidth, employeeStartY + 75)

  // Detalles laborales
  const laborDetailsStartY = employeeStartY + 100
  doc.text('Días Laborados: 30', margin, laborDetailsStartY)
  doc.text('Días No Laborados: 0', margin + columnWidth / 2, laborDetailsStartY)
  doc.text('Días Subsidiados: 0', margin + columnWidth, laborDetailsStartY)
  doc.text('Periodo Vacacional: 0', margin, laborDetailsStartY + 15)

  // Detalle de Ingresos
  const ingresosStartY = laborDetailsStartY + 40
  doc.autoTable({
    startY: ingresosStartY,
    head: [['CÓDIGO', 'CONCEPTO', 'MONTO']],
    body: [
      ['0001', 'Rem. Básica', '1500.00'],
      ['0002', 'Asignación Familiar', '100.00'],
      ['0003', 'Horas Extras', '200.00']
    ],
    theme: 'grid',
    styles: { halign: 'center' },
    headStyles: { fillColor: [0, 0, 0] }
  })
  const ingresosEndY = doc.lastAutoTable.finalY
  doc.text('TOTAL INGRESOS', margin, ingresosEndY + 10)
  doc.text('1800.00', pageWidth - margin, ingresosEndY + 10, { align: 'right' })

  // Detalle de Descuentos
  doc.autoTable({
    startY: ingresosEndY + 30,
    head: [['CÓDIGO', 'CONCEPTO', 'MONTO']],
    body: [
      ['0035', 'AFP - Aporte Obligatorio', '150.00'],
      ['0039', 'AFP - Comisión', '20.00'],
      ['0045', 'Seguro de Salud', '30.00']
    ],
    theme: 'grid',
    styles: { halign: 'center' },
    headStyles: { fillColor: [0, 0, 0] }
  })
  const descuentosEndY = doc.lastAutoTable.finalY
  doc.text('TOTAL DESCUENTOS', margin, descuentosEndY + 10)
  doc.text('200.00', pageWidth - margin, descuentosEndY + 10, { align: 'right' })

  // Detalle de Aportes
  doc.autoTable({
    startY: descuentosEndY + 30,
    head: [['CÓDIGO', 'CONCEPTO', 'MONTO']],
    body: [
      ['075', 'ESSALUD', '116.55'],
      ['076', 'EPS', '116.55']
    ],
    theme: 'grid',
    styles: { halign: 'center' },
    headStyles: { fillColor: [0, 0, 0] }
  })
  const aportesEndY = doc.lastAutoTable.finalY
  doc.text('TOTAL APORTES', margin, aportesEndY + 10)
  doc.text('233.10', pageWidth - margin, aportesEndY + 10, { align: 'right' })

  // Resumen financiero
  doc.text('MONTO IMPONIBLE', margin, aportesEndY + 30)
  doc.text('400.00', pageWidth - margin, aportesEndY + 30, { align: 'right' })
  doc.text('NETO A PAGAR', margin, aportesEndY + 50)
  doc.text('1,566.90', pageWidth - margin, aportesEndY + 50, { align: 'right' })
  doc.text(
    'NETO A PAGAR (En letras): Mil quinientos sesenta y seis con 90/100 Soles',
    margin,
    aportesEndY + 70
  )

  // Observaciones
  doc.text('OBSERVACIÓN', margin, aportesEndY + 100)

  // Pie de página
  doc.text('Código QR', margin, pageHeight - 80)
  doc.text('Este documento es auténtico...', pageWidth / 2, pageHeight - 80, { align: 'center' })
  doc.text('https://mef.gob.pe/boletapago', pageWidth - margin, pageHeight - 80, { align: 'right' })

  // Guardar PDF
  doc.output('dataurlnewwindow')
  // doc.save('boleta_pago.pdf')
}
