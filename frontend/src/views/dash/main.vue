<template>
  <div class="page">
    <div class="px-4 pt-2">
      <div class="max-w-7xl mx-auto">
        <div class="space-y-1">
          <span class="text-sm font-medium"> Resumen de los indicadores </span>
          <h2 class="font-bold text-gray tracking-tight">Dashboard</h2>
        </div>
      </div>
    </div>

    <div class="page-body">
      <div class="container-xl">
        <div class="row row-deck row-cards">
          <div class="col-12">
            <div class="row row-cards">
              <div class="col-sm-6 col-lg-3">
                <div class="card card-sm">
                  <div class="card-body">
                    <div class="row align-items-center">
                      <div class="col-auto">
                        <span class="bg-primary text-white avatar">
                          <!-- Download SVG icon from http://tabler.io/icons/icon/currency-dollar -->
                          <svg
                            xmlns="http://www.w3.org/2000/svg"
                            width="24"
                            height="24"
                            viewBox="0 0 24 24"
                            fill="none"
                            stroke="currentColor"
                            stroke-width="2"
                            stroke-linecap="round"
                            stroke-linejoin="round"
                            class="icon icon-1"
                          >
                            <path
                              d="M16.7 8a3 3 0 0 0 -2.7 -2h-4a3 3 0 0 0 0 6h4a3 3 0 0 1 0 6h-4a3 3 0 0 1 -2.7 -2"
                            ></path>
                            <path d="M12 3v3m0 12v3"></path>
                          </svg>
                        </span>
                      </div>
                      <div class="col">
                        <div class="font-weight-medium">132 Sales</div>
                        <div class="text-secondary">12 waiting payments</div>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
              <div class="col-sm-6 col-lg-3">
                <div class="card card-sm">
                  <div class="card-body">
                    <div class="row align-items-center">
                      <div class="col-auto">
                        <span class="bg-green text-white avatar">
                          <!-- Download SVG icon from http://tabler.io/icons/icon/shopping-cart -->
                          <svg
                            xmlns="http://www.w3.org/2000/svg"
                            width="24"
                            height="24"
                            viewBox="0 0 24 24"
                            fill="none"
                            stroke="currentColor"
                            stroke-width="2"
                            stroke-linecap="round"
                            stroke-linejoin="round"
                            class="icon icon-1"
                          >
                            <path d="M6 19m-2 0a2 2 0 1 0 4 0a2 2 0 1 0 -4 0"></path>
                            <path d="M17 19m-2 0a2 2 0 1 0 4 0a2 2 0 1 0 -4 0"></path>
                            <path d="M17 17h-11v-14h-2"></path>
                            <path d="M6 5l14 1l-1 7h-13"></path>
                          </svg>
                        </span>
                      </div>
                      <div class="col">
                        <div class="font-weight-medium">78 Orders</div>
                        <div class="text-secondary">32 shipped</div>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
              <div class="col-sm-6 col-lg-3">
                <div class="card card-sm">
                  <div class="card-body">
                    <div class="row align-items-center">
                      <div class="col-auto">
                        <span class="bg-x text-white avatar">
                          <!-- Download SVG icon from http://tabler.io/icons/icon/brand-x -->
                          <svg
                            xmlns="http://www.w3.org/2000/svg"
                            width="24"
                            height="24"
                            viewBox="0 0 24 24"
                            fill="none"
                            stroke="currentColor"
                            stroke-width="2"
                            stroke-linecap="round"
                            stroke-linejoin="round"
                            class="icon icon-1"
                          >
                            <path d="M4 4l11.733 16h4.267l-11.733 -16z"></path>
                            <path d="M4 20l6.768 -6.768m2.46 -2.46l6.772 -6.772"></path>
                          </svg>
                        </span>
                      </div>
                      <div class="col">
                        <div class="font-weight-medium">623 Shares</div>
                        <div class="text-secondary">16 today</div>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
              <div class="col-sm-6 col-lg-3">
                <div class="card card-sm">
                  <div class="card-body">
                    <div class="row align-items-center">
                      <div class="col-auto">
                        <span class="bg-facebook text-white avatar">
                          <!-- Download SVG icon from http://tabler.io/icons/icon/brand-facebook -->
                          <svg
                            xmlns="http://www.w3.org/2000/svg"
                            width="24"
                            height="24"
                            viewBox="0 0 24 24"
                            fill="none"
                            stroke="currentColor"
                            stroke-width="2"
                            stroke-linecap="round"
                            stroke-linejoin="round"
                            class="icon icon-1"
                          >
                            <path
                              d="M7 10v4h3v7h4v-7h3l1 -4h-4v-2a1 1 0 0 1 1 -1h3v-4h-3a5 5 0 0 0 -5 5v2h-3"
                            ></path>
                          </svg>
                        </span>
                      </div>
                      <div class="col">
                        <div class="font-weight-medium">132 Likes</div>
                        <div class="text-secondary">21 today</div>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
          <div class="col-md-6 col-lg-4">
            <div class="card">
              <div class="card-header">
                <h3 class="card-title">Social Media Traffic</h3>
              </div>
              <table class="table card-table table-vcenter">
                <thead>
                  <tr>
                    <th>Network</th>
                    <th colspan="2">Visitors</th>
                  </tr>
                </thead>
                <tbody>
                  <tr>
                    <td>Instagram</td>
                    <td>3,550</td>
                    <td class="w-50">
                      <div class="progress progress-xs">
                        <div class="progress-bar bg-primary" style="width: 71%"></div>
                      </div>
                    </td>
                  </tr>
                  <tr>
                    <td>Twitter</td>
                    <td>1,798</td>
                    <td class="w-50">
                      <div class="progress progress-xs">
                        <div class="progress-bar bg-primary" style="width: 35.96%"></div>
                      </div>
                    </td>
                  </tr>
                  <tr>
                    <td>Facebook</td>
                    <td>1,245</td>
                    <td class="w-50">
                      <div class="progress progress-xs">
                        <div class="progress-bar bg-primary" style="width: 24.9%"></div>
                      </div>
                    </td>
                  </tr>
                  <tr>
                    <td>TikTok</td>
                    <td>986</td>
                    <td class="w-50">
                      <div class="progress progress-xs">
                        <div class="progress-bar bg-primary" style="width: 19.72%"></div>
                      </div>
                    </td>
                  </tr>
                  <tr>
                    <td>Pinterest</td>
                    <td>854</td>
                    <td class="w-50">
                      <div class="progress progress-xs">
                        <div
                          class="progress-bar bg-primary"
                          style="width: 17.080000000000002%"
                        ></div>
                      </div>
                    </td>
                  </tr>
                  <tr>
                    <td>VK</td>
                    <td>650</td>
                    <td class="w-50">
                      <div class="progress progress-xs">
                        <div class="progress-bar bg-primary" style="width: 13%"></div>
                      </div>
                    </td>
                  </tr>
                  <tr>
                    <td>Pinterest</td>
                    <td>420</td>
                    <td class="w-50">
                      <div class="progress progress-xs">
                        <div class="progress-bar bg-primary" style="width: 8.4%"></div>
                      </div>
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>
          </div>
          <div class="col-md-12 col-lg-8">
            <div class="card">
              <div class="card-header">
                <h3 class="card-title">Tasks</h3>
              </div>
              <div class="table-responsive">
                <table class="table card-table table-vcenter">
                  <tbody>
                    <tr>
                      <td class="w-1 pe-0">
                        <input
                          type="checkbox"
                          class="form-check-input m-0 align-middle"
                          aria-label="Select task"
                        />
                      </td>
                      <td class="w-100">
                        <a href="#" class="text-reset">Extend the data model.</a>
                      </td>
                      <td class="text-nowrap text-secondary">
                        <!-- Download SVG icon from http://tabler.io/icons/icon/calendar -->
                        <svg
                          xmlns="http://www.w3.org/2000/svg"
                          width="24"
                          height="24"
                          viewBox="0 0 24 24"
                          fill="none"
                          stroke="currentColor"
                          stroke-width="2"
                          stroke-linecap="round"
                          stroke-linejoin="round"
                          class="icon icon-1"
                        >
                          <path
                            d="M4 7a2 2 0 0 1 2 -2h12a2 2 0 0 1 2 2v12a2 2 0 0 1 -2 2h-12a2 2 0 0 1 -2 -2v-12z"
                          ></path>
                          <path d="M16 3v4"></path>
                          <path d="M8 3v4"></path>
                          <path d="M4 11h16"></path>
                          <path d="M11 15h1"></path>
                          <path d="M12 15v3"></path>
                        </svg>
                        December 11, 2024
                      </td>
                      <td class="text-nowrap">
                        <a href="#" class="text-secondary">
                          <!-- Download SVG icon from http://tabler.io/icons/icon/check -->
                          <svg
                            xmlns="http://www.w3.org/2000/svg"
                            width="24"
                            height="24"
                            viewBox="0 0 24 24"
                            fill="none"
                            stroke="currentColor"
                            stroke-width="2"
                            stroke-linecap="round"
                            stroke-linejoin="round"
                            class="icon icon-1"
                          >
                            <path d="M5 12l5 5l10 -10"></path>
                          </svg>
                          2/7
                        </a>
                      </td>
                      <td class="text-nowrap">
                        <a href="#" class="text-secondary">
                          <!-- Download SVG icon from http://tabler.io/icons/icon/message -->
                          <svg
                            xmlns="http://www.w3.org/2000/svg"
                            width="24"
                            height="24"
                            viewBox="0 0 24 24"
                            fill="none"
                            stroke="currentColor"
                            stroke-width="2"
                            stroke-linecap="round"
                            stroke-linejoin="round"
                            class="icon icon-1"
                          >
                            <path d="M8 9h8"></path>
                            <path d="M8 13h6"></path>
                            <path
                              d="M18 4a3 3 0 0 1 3 3v8a3 3 0 0 1 -3 3h-5l-5 3v-3h-2a3 3 0 0 1 -3 -3v-8a3 3 0 0 1 3 -3h12z"
                            ></path>
                          </svg>
                          3</a
                        >
                      </td>
                      <td>
                        <span
                          class="avatar avatar-sm"
                          style="background-image: url(./static/avatars/000m.jpg)"
                        ></span>
                      </td>
                    </tr>
                    <tr>
                      <td class="w-1 pe-0">
                        <input
                          type="checkbox"
                          class="form-check-input m-0 align-middle"
                          aria-label="Select task"
                        />
                      </td>
                      <td class="w-100">
                        <a href="#" class="text-reset">Verify the event flow.</a>
                      </td>
                      <td class="text-nowrap text-secondary">
                        <!-- Download SVG icon from http://tabler.io/icons/icon/calendar -->
                        <svg
                          xmlns="http://www.w3.org/2000/svg"
                          width="24"
                          height="24"
                          viewBox="0 0 24 24"
                          fill="none"
                          stroke="currentColor"
                          stroke-width="2"
                          stroke-linecap="round"
                          stroke-linejoin="round"
                          class="icon icon-1"
                        >
                          <path
                            d="M4 7a2 2 0 0 1 2 -2h12a2 2 0 0 1 2 2v12a2 2 0 0 1 -2 2h-12a2 2 0 0 1 -2 -2v-12z"
                          ></path>
                          <path d="M16 3v4"></path>
                          <path d="M8 3v4"></path>
                          <path d="M4 11h16"></path>
                          <path d="M11 15h1"></path>
                          <path d="M12 15v3"></path>
                        </svg>
                        October 20, 2024
                      </td>
                      <td class="text-nowrap">
                        <a href="#" class="text-secondary">
                          <!-- Download SVG icon from http://tabler.io/icons/icon/check -->
                          <svg
                            xmlns="http://www.w3.org/2000/svg"
                            width="24"
                            height="24"
                            viewBox="0 0 24 24"
                            fill="none"
                            stroke="currentColor"
                            stroke-width="2"
                            stroke-linecap="round"
                            stroke-linejoin="round"
                            class="icon icon-1"
                          >
                            <path d="M5 12l5 5l10 -10"></path>
                          </svg>
                          0/5
                        </a>
                      </td>
                      <td class="text-nowrap">
                        <a href="#" class="text-secondary">
                          <!-- Download SVG icon from http://tabler.io/icons/icon/message -->
                          <svg
                            xmlns="http://www.w3.org/2000/svg"
                            width="24"
                            height="24"
                            viewBox="0 0 24 24"
                            fill="none"
                            stroke="currentColor"
                            stroke-width="2"
                            stroke-linecap="round"
                            stroke-linejoin="round"
                            class="icon icon-1"
                          >
                            <path d="M8 9h8"></path>
                            <path d="M8 13h6"></path>
                            <path
                              d="M18 4a3 3 0 0 1 3 3v8a3 3 0 0 1 -3 3h-5l-5 3v-3h-2a3 3 0 0 1 -3 -3v-8a3 3 0 0 1 3 -3h12z"
                            ></path>
                          </svg>
                          0</a
                        >
                      </td>
                      <td>
                        <span class="avatar avatar-sm">JL</span>
                      </td>
                    </tr>
                    <tr>
                      <td class="w-1 pe-0">
                        <input
                          type="checkbox"
                          class="form-check-input m-0 align-middle"
                          aria-label="Select task"
                        />
                      </td>
                      <td class="w-100">
                        <a href="#" class="text-reset">Database backup and maintenance</a>
                      </td>
                      <td class="text-nowrap text-secondary">
                        <!-- Download SVG icon from http://tabler.io/icons/icon/calendar -->
                        <svg
                          xmlns="http://www.w3.org/2000/svg"
                          width="24"
                          height="24"
                          viewBox="0 0 24 24"
                          fill="none"
                          stroke="currentColor"
                          stroke-width="2"
                          stroke-linecap="round"
                          stroke-linejoin="round"
                          class="icon icon-1"
                        >
                          <path
                            d="M4 7a2 2 0 0 1 2 -2h12a2 2 0 0 1 2 2v12a2 2 0 0 1 -2 2h-12a2 2 0 0 1 -2 -2v-12z"
                          ></path>
                          <path d="M16 3v4"></path>
                          <path d="M8 3v4"></path>
                          <path d="M4 11h16"></path>
                          <path d="M11 15h1"></path>
                          <path d="M12 15v3"></path>
                        </svg>
                        October 20, 2024
                      </td>
                      <td class="text-nowrap">
                        <a href="#" class="text-secondary">
                          <!-- Download SVG icon from http://tabler.io/icons/icon/check -->
                          <svg
                            xmlns="http://www.w3.org/2000/svg"
                            width="24"
                            height="24"
                            viewBox="0 0 24 24"
                            fill="none"
                            stroke="currentColor"
                            stroke-width="2"
                            stroke-linecap="round"
                            stroke-linejoin="round"
                            class="icon icon-1"
                          >
                            <path d="M5 12l5 5l10 -10"></path>
                          </svg>
                          0/5
                        </a>
                      </td>
                      <td class="text-nowrap">
                        <a href="#" class="text-secondary">
                          <!-- Download SVG icon from http://tabler.io/icons/icon/message -->
                          <svg
                            xmlns="http://www.w3.org/2000/svg"
                            width="24"
                            height="24"
                            viewBox="0 0 24 24"
                            fill="none"
                            stroke="currentColor"
                            stroke-width="2"
                            stroke-linecap="round"
                            stroke-linejoin="round"
                            class="icon icon-1"
                          >
                            <path d="M8 9h8"></path>
                            <path d="M8 13h6"></path>
                            <path
                              d="M18 4a3 3 0 0 1 3 3v8a3 3 0 0 1 -3 3h-5l-5 3v-3h-2a3 3 0 0 1 -3 -3v-8a3 3 0 0 1 3 -3h12z"
                            ></path>
                          </svg>
                          0</a
                        >
                      </td>
                      <td>
                        <span
                          class="avatar avatar-sm"
                          style="background-image: url(./static/avatars/002m.jpg)"
                        ></span>
                      </td>
                    </tr>
                    <tr>
                      <td class="w-1 pe-0">
                        <input
                          type="checkbox"
                          class="form-check-input m-0 align-middle"
                          aria-label="Select task"
                        />
                      </td>
                      <td class="w-100">
                        <a href="#" class="text-reset">Identify the implementation team.</a>
                      </td>
                      <td class="text-nowrap text-secondary">
                        <!-- Download SVG icon from http://tabler.io/icons/icon/calendar -->
                        <svg
                          xmlns="http://www.w3.org/2000/svg"
                          width="24"
                          height="24"
                          viewBox="0 0 24 24"
                          fill="none"
                          stroke="currentColor"
                          stroke-width="2"
                          stroke-linecap="round"
                          stroke-linejoin="round"
                          class="icon icon-1"
                        >
                          <path
                            d="M4 7a2 2 0 0 1 2 -2h12a2 2 0 0 1 2 2v12a2 2 0 0 1 -2 2h-12a2 2 0 0 1 -2 -2v-12z"
                          ></path>
                          <path d="M16 3v4"></path>
                          <path d="M8 3v4"></path>
                          <path d="M4 11h16"></path>
                          <path d="M11 15h1"></path>
                          <path d="M12 15v3"></path>
                        </svg>
                        January 14, 2025
                      </td>
                      <td class="text-nowrap">
                        <a href="#" class="text-secondary">
                          <!-- Download SVG icon from http://tabler.io/icons/icon/check -->
                          <svg
                            xmlns="http://www.w3.org/2000/svg"
                            width="24"
                            height="24"
                            viewBox="0 0 24 24"
                            fill="none"
                            stroke="currentColor"
                            stroke-width="2"
                            stroke-linecap="round"
                            stroke-linejoin="round"
                            class="icon icon-1"
                          >
                            <path d="M5 12l5 5l10 -10"></path>
                          </svg>
                          6/10
                        </a>
                      </td>
                      <td class="text-nowrap">
                        <a href="#" class="text-secondary">
                          <!-- Download SVG icon from http://tabler.io/icons/icon/message -->
                          <svg
                            xmlns="http://www.w3.org/2000/svg"
                            width="24"
                            height="24"
                            viewBox="0 0 24 24"
                            fill="none"
                            stroke="currentColor"
                            stroke-width="2"
                            stroke-linecap="round"
                            stroke-linejoin="round"
                            class="icon icon-1"
                          >
                            <path d="M8 9h8"></path>
                            <path d="M8 13h6"></path>
                            <path
                              d="M18 4a3 3 0 0 1 3 3v8a3 3 0 0 1 -3 3h-5l-5 3v-3h-2a3 3 0 0 1 -3 -3v-8a3 3 0 0 1 3 -3h12z"
                            ></path>
                          </svg>
                          12</a
                        >
                      </td>
                      <td>
                        <span
                          class="avatar avatar-sm"
                          style="background-image: url(./static/avatars/003m.jpg)"
                        ></span>
                      </td>
                    </tr>
                    <tr>
                      <td class="w-1 pe-0">
                        <input
                          type="checkbox"
                          class="form-check-input m-0 align-middle"
                          aria-label="Select task"
                        />
                      </td>
                      <td class="w-100">
                        <a href="#" class="text-reset">Define users and workflow</a>
                      </td>
                      <td class="text-nowrap text-secondary">
                        <!-- Download SVG icon from http://tabler.io/icons/icon/calendar -->
                        <svg
                          xmlns="http://www.w3.org/2000/svg"
                          width="24"
                          height="24"
                          viewBox="0 0 24 24"
                          fill="none"
                          stroke="currentColor"
                          stroke-width="2"
                          stroke-linecap="round"
                          stroke-linejoin="round"
                          class="icon icon-1"
                        >
                          <path
                            d="M4 7a2 2 0 0 1 2 -2h12a2 2 0 0 1 2 2v12a2 2 0 0 1 -2 2h-12a2 2 0 0 1 -2 -2v-12z"
                          ></path>
                          <path d="M16 3v4"></path>
                          <path d="M8 3v4"></path>
                          <path d="M4 11h16"></path>
                          <path d="M11 15h1"></path>
                          <path d="M12 15v3"></path>
                        </svg>
                        October 20, 2024
                      </td>
                      <td class="text-nowrap">
                        <a href="#" class="text-secondary">
                          <!-- Download SVG icon from http://tabler.io/icons/icon/check -->
                          <svg
                            xmlns="http://www.w3.org/2000/svg"
                            width="24"
                            height="24"
                            viewBox="0 0 24 24"
                            fill="none"
                            stroke="currentColor"
                            stroke-width="2"
                            stroke-linecap="round"
                            stroke-linejoin="round"
                            class="icon icon-1"
                          >
                            <path d="M5 12l5 5l10 -10"></path>
                          </svg>
                          0/5
                        </a>
                      </td>
                      <td class="text-nowrap">
                        <a href="#" class="text-secondary">
                          <!-- Download SVG icon from http://tabler.io/icons/icon/message -->
                          <svg
                            xmlns="http://www.w3.org/2000/svg"
                            width="24"
                            height="24"
                            viewBox="0 0 24 24"
                            fill="none"
                            stroke="currentColor"
                            stroke-width="2"
                            stroke-linecap="round"
                            stroke-linejoin="round"
                            class="icon icon-1"
                          >
                            <path d="M8 9h8"></path>
                            <path d="M8 13h6"></path>
                            <path
                              d="M18 4a3 3 0 0 1 3 3v8a3 3 0 0 1 -3 3h-5l-5 3v-3h-2a3 3 0 0 1 -3 -3v-8a3 3 0 0 1 3 -3h12z"
                            ></path>
                          </svg>
                          0</a
                        >
                      </td>
                      <td>
                        <span
                          class="avatar avatar-sm"
                          style="background-image: url(./static/avatars/000f.jpg)"
                        ></span>
                      </td>
                    </tr>
                    <tr>
                      <td class="w-1 pe-0">
                        <input
                          type="checkbox"
                          class="form-check-input m-0 align-middle"
                          aria-label="Select task"
                        />
                      </td>
                      <td class="w-100">
                        <a href="#" class="text-reset">Check Pull Requests</a>
                      </td>
                      <td class="text-nowrap text-secondary">
                        <!-- Download SVG icon from http://tabler.io/icons/icon/calendar -->
                        <svg
                          xmlns="http://www.w3.org/2000/svg"
                          width="24"
                          height="24"
                          viewBox="0 0 24 24"
                          fill="none"
                          stroke="currentColor"
                          stroke-width="2"
                          stroke-linecap="round"
                          stroke-linejoin="round"
                          class="icon icon-1"
                        >
                          <path
                            d="M4 7a2 2 0 0 1 2 -2h12a2 2 0 0 1 2 2v12a2 2 0 0 1 -2 2h-12a2 2 0 0 1 -2 -2v-12z"
                          ></path>
                          <path d="M16 3v4"></path>
                          <path d="M8 3v4"></path>
                          <path d="M4 11h16"></path>
                          <path d="M11 15h1"></path>
                          <path d="M12 15v3"></path>
                        </svg>
                        January 16, 2025
                      </td>
                      <td class="text-nowrap">
                        <a href="#" class="text-secondary">
                          <!-- Download SVG icon from http://tabler.io/icons/icon/check -->
                          <svg
                            xmlns="http://www.w3.org/2000/svg"
                            width="24"
                            height="24"
                            viewBox="0 0 24 24"
                            fill="none"
                            stroke="currentColor"
                            stroke-width="2"
                            stroke-linecap="round"
                            stroke-linejoin="round"
                            class="icon icon-1"
                          >
                            <path d="M5 12l5 5l10 -10"></path>
                          </svg>
                          2/9
                        </a>
                      </td>
                      <td class="text-nowrap">
                        <a href="#" class="text-secondary">
                          <!-- Download SVG icon from http://tabler.io/icons/icon/message -->
                          <svg
                            xmlns="http://www.w3.org/2000/svg"
                            width="24"
                            height="24"
                            viewBox="0 0 24 24"
                            fill="none"
                            stroke="currentColor"
                            stroke-width="2"
                            stroke-linecap="round"
                            stroke-linejoin="round"
                            class="icon icon-1"
                          >
                            <path d="M8 9h8"></path>
                            <path d="M8 13h6"></path>
                            <path
                              d="M18 4a3 3 0 0 1 3 3v8a3 3 0 0 1 -3 3h-5l-5 3v-3h-2a3 3 0 0 1 -3 -3v-8a3 3 0 0 1 3 -3h12z"
                            ></path>
                          </svg>
                          3</a
                        >
                      </td>
                      <td>
                        <span
                          class="avatar avatar-sm"
                          style="background-image: url(./static/avatars/001f.jpg)"
                        ></span>
                      </td>
                    </tr>
                  </tbody>
                </table>
              </div>
            </div>
          </div>
          <div class="col-12">
            <div class="card">
              <div class="card-header">
                <h3 class="card-title">Invoices</h3>
              </div>
              <div class="card-body border-bottom py-3">
                <div class="d-flex">
                  <div class="text-secondary">
                    Show
                    <div class="mx-2 d-inline-block">
                      <input
                        type="text"
                        class="form-control form-control-sm"
                        value="8"
                        size="3"
                        aria-label="Invoices count"
                      />
                    </div>
                    entries
                  </div>
                  <div class="ms-auto text-secondary">
                    Search:
                    <div class="ms-2 d-inline-block">
                      <input
                        type="text"
                        class="form-control form-control-sm"
                        aria-label="Search invoice"
                      />
                    </div>
                  </div>
                </div>
              </div>
              <div class="table-responsive">
                <table class="table card-table table-vcenter text-nowrap datatable">
                  <thead>
                    <tr>
                      <th class="w-1">
                        <input
                          class="form-check-input m-0 align-middle"
                          type="checkbox"
                          aria-label="Select all invoices"
                        />
                      </th>
                      <th class="w-1">
                        No.
                        <!-- Download SVG icon from http://tabler.io/icons/icon/chevron-up -->
                        <svg
                          xmlns="http://www.w3.org/2000/svg"
                          width="24"
                          height="24"
                          viewBox="0 0 24 24"
                          fill="none"
                          stroke="currentColor"
                          stroke-width="2"
                          stroke-linecap="round"
                          stroke-linejoin="round"
                          class="icon icon-sm icon-thick icon-2"
                        >
                          <path d="M6 15l6 -6l6 6"></path>
                        </svg>
                      </th>
                      <th>Invoice Subject</th>
                      <th>Client</th>
                      <th>VAT No.</th>
                      <th>Created</th>
                      <th>Status</th>
                      <th>Price</th>
                      <th></th>
                    </tr>
                  </thead>
                  <tbody>
                    <tr>
                      <td>
                        <input
                          class="form-check-input m-0 align-middle"
                          type="checkbox"
                          aria-label="Select invoice"
                        />
                      </td>
                      <td><span class="text-secondary">001401</span></td>
                      <td>
                        <a href="invoice.html" class="text-reset" tabindex="-1">Design Works</a>
                      </td>
                      <td>
                        <span class="flag flag-xs flag-country-us me-2"></span>
                        Carlson Limited
                      </td>
                      <td>87956621</td>
                      <td>15 Dec 2017</td>
                      <td><span class="badge bg-success me-1"></span> Paid</td>
                      <td>$887</td>
                      <td class="text-end">
                        <span class="dropdown">
                          <button
                            class="btn dropdown-toggle align-text-top"
                            data-bs-boundary="viewport"
                            data-bs-toggle="dropdown"
                          >
                            Actions
                          </button>
                          <div class="dropdown-menu dropdown-menu-end">
                            <a class="dropdown-item" href="#"> Action </a>
                            <a class="dropdown-item" href="#"> Another action </a>
                          </div>
                        </span>
                      </td>
                    </tr>
                    <tr>
                      <td>
                        <input
                          class="form-check-input m-0 align-middle"
                          type="checkbox"
                          aria-label="Select invoice"
                        />
                      </td>
                      <td><span class="text-secondary">001402</span></td>
                      <td>
                        <a href="invoice.html" class="text-reset" tabindex="-1">UX Wireframes</a>
                      </td>
                      <td>
                        <span class="flag flag-xs flag-country-gb me-2"></span>
                        Adobe
                      </td>
                      <td>87956421</td>
                      <td>12 Apr 2017</td>
                      <td><span class="badge bg-warning me-1"></span> Pending</td>
                      <td>$1200</td>
                      <td class="text-end">
                        <span class="dropdown">
                          <button
                            class="btn dropdown-toggle align-text-top"
                            data-bs-boundary="viewport"
                            data-bs-toggle="dropdown"
                          >
                            Actions
                          </button>
                          <div class="dropdown-menu dropdown-menu-end">
                            <a class="dropdown-item" href="#"> Action </a>
                            <a class="dropdown-item" href="#"> Another action </a>
                          </div>
                        </span>
                      </td>
                    </tr>
                    <tr>
                      <td>
                        <input
                          class="form-check-input m-0 align-middle"
                          type="checkbox"
                          aria-label="Select invoice"
                        />
                      </td>
                      <td><span class="text-secondary">001403</span></td>
                      <td>
                        <a href="invoice.html" class="text-reset" tabindex="-1">New Dashboard</a>
                      </td>
                      <td>
                        <span class="flag flag-xs flag-country-de me-2"></span>
                        Bluewolf
                      </td>
                      <td>87952621</td>
                      <td>23 Oct 2017</td>
                      <td><span class="badge bg-warning me-1"></span> Pending</td>
                      <td>$534</td>
                      <td class="text-end">
                        <span class="dropdown">
                          <button
                            class="btn dropdown-toggle align-text-top"
                            data-bs-boundary="viewport"
                            data-bs-toggle="dropdown"
                          >
                            Actions
                          </button>
                          <div class="dropdown-menu dropdown-menu-end">
                            <a class="dropdown-item" href="#"> Action </a>
                            <a class="dropdown-item" href="#"> Another action </a>
                          </div>
                        </span>
                      </td>
                    </tr>
                    <tr>
                      <td>
                        <input
                          class="form-check-input m-0 align-middle"
                          type="checkbox"
                          aria-label="Select invoice"
                        />
                      </td>
                      <td><span class="text-secondary">001404</span></td>
                      <td>
                        <a href="invoice.html" class="text-reset" tabindex="-1">Landing Page</a>
                      </td>
                      <td>
                        <span class="flag flag-xs flag-country-br me-2"></span>
                        Salesforce
                      </td>
                      <td>87953421</td>
                      <td>2 Sep 2017</td>
                      <td><span class="badge bg-secondary me-1"></span> Due in 2 Weeks</td>
                      <td>$1500</td>
                      <td class="text-end">
                        <span class="dropdown">
                          <button
                            class="btn dropdown-toggle align-text-top"
                            data-bs-boundary="viewport"
                            data-bs-toggle="dropdown"
                          >
                            Actions
                          </button>
                          <div class="dropdown-menu dropdown-menu-end">
                            <a class="dropdown-item" href="#"> Action </a>
                            <a class="dropdown-item" href="#"> Another action </a>
                          </div>
                        </span>
                      </td>
                    </tr>
                    <tr>
                      <td>
                        <input
                          class="form-check-input m-0 align-middle"
                          type="checkbox"
                          aria-label="Select invoice"
                        />
                      </td>
                      <td><span class="text-secondary">001405</span></td>
                      <td>
                        <a href="invoice.html" class="text-reset" tabindex="-1"
                          >Marketing Templates</a
                        >
                      </td>
                      <td>
                        <span class="flag flag-xs flag-country-pl me-2"></span>
                        Printic
                      </td>
                      <td>87956621</td>
                      <td>29 Jan 2018</td>
                      <td><span class="badge bg-danger me-1"></span> Paid Today</td>
                      <td>$648</td>
                      <td class="text-end">
                        <span class="dropdown">
                          <button
                            class="btn dropdown-toggle align-text-top"
                            data-bs-boundary="viewport"
                            data-bs-toggle="dropdown"
                          >
                            Actions
                          </button>
                          <div class="dropdown-menu dropdown-menu-end">
                            <a class="dropdown-item" href="#"> Action </a>
                            <a class="dropdown-item" href="#"> Another action </a>
                          </div>
                        </span>
                      </td>
                    </tr>
                    <tr>
                      <td>
                        <input
                          class="form-check-input m-0 align-middle"
                          type="checkbox"
                          aria-label="Select invoice"
                        />
                      </td>
                      <td><span class="text-secondary">001406</span></td>
                      <td>
                        <a href="invoice.html" class="text-reset" tabindex="-1"
                          >Sales Presentation</a
                        >
                      </td>
                      <td>
                        <span class="flag flag-xs flag-country-br me-2"></span>
                        Tabdaq
                      </td>
                      <td>87956621</td>
                      <td>4 Feb 2018</td>
                      <td><span class="badge bg-secondary me-1"></span> Due in 3 Weeks</td>
                      <td>$300</td>
                      <td class="text-end">
                        <span class="dropdown">
                          <button
                            class="btn dropdown-toggle align-text-top"
                            data-bs-boundary="viewport"
                            data-bs-toggle="dropdown"
                          >
                            Actions
                          </button>
                          <div class="dropdown-menu dropdown-menu-end">
                            <a class="dropdown-item" href="#"> Action </a>
                            <a class="dropdown-item" href="#"> Another action </a>
                          </div>
                        </span>
                      </td>
                    </tr>
                    <tr>
                      <td>
                        <input
                          class="form-check-input m-0 align-middle"
                          type="checkbox"
                          aria-label="Select invoice"
                        />
                      </td>
                      <td><span class="text-secondary">001407</span></td>
                      <td>
                        <a href="invoice.html" class="text-reset" tabindex="-1">Logo &amp; Print</a>
                      </td>
                      <td>
                        <span class="flag flag-xs flag-country-us me-2"></span>
                        Apple
                      </td>
                      <td>87956621</td>
                      <td>22 Mar 2018</td>
                      <td><span class="badge bg-success me-1"></span> Paid Today</td>
                      <td>$2500</td>
                      <td class="text-end">
                        <span class="dropdown">
                          <button
                            class="btn dropdown-toggle align-text-top"
                            data-bs-boundary="viewport"
                            data-bs-toggle="dropdown"
                          >
                            Actions
                          </button>
                          <div class="dropdown-menu dropdown-menu-end">
                            <a class="dropdown-item" href="#"> Action </a>
                            <a class="dropdown-item" href="#"> Another action </a>
                          </div>
                        </span>
                      </td>
                    </tr>
                    <tr>
                      <td>
                        <input
                          class="form-check-input m-0 align-middle"
                          type="checkbox"
                          aria-label="Select invoice"
                        />
                      </td>
                      <td><span class="text-secondary">001408</span></td>
                      <td><a href="invoice.html" class="text-reset" tabindex="-1">Icons</a></td>
                      <td>
                        <span class="flag flag-xs flag-country-pl me-2"></span>
                        Tookapic
                      </td>
                      <td>87956621</td>
                      <td>13 May 2018</td>
                      <td><span class="badge bg-success me-1"></span> Paid Today</td>
                      <td>$940</td>
                      <td class="text-end">
                        <span class="dropdown">
                          <button
                            class="btn dropdown-toggle align-text-top"
                            data-bs-boundary="viewport"
                            data-bs-toggle="dropdown"
                          >
                            Actions
                          </button>
                          <div class="dropdown-menu dropdown-menu-end">
                            <a class="dropdown-item" href="#"> Action </a>
                            <a class="dropdown-item" href="#"> Another action </a>
                          </div>
                        </span>
                      </td>
                    </tr>
                  </tbody>
                </table>
              </div>
              <div class="card-footer d-flex align-items-center">
                <p class="m-0 text-secondary">
                  Showing <span>1</span> to <span>8</span> of <span>16</span> entries
                </p>
                <ul class="pagination m-0 ms-auto">
                  <li class="page-item disabled">
                    <a class="page-link" href="#" tabindex="-1" aria-disabled="true">
                      <!-- Download SVG icon from http://tabler.io/icons/icon/chevron-left -->
                      <svg
                        xmlns="http://www.w3.org/2000/svg"
                        width="24"
                        height="24"
                        viewBox="0 0 24 24"
                        fill="none"
                        stroke="currentColor"
                        stroke-width="2"
                        stroke-linecap="round"
                        stroke-linejoin="round"
                        class="icon icon-1"
                      >
                        <path d="M15 6l-6 6l6 6"></path>
                      </svg>
                      prev
                    </a>
                  </li>
                  <li class="page-item"><a class="page-link" href="#">1</a></li>
                  <li class="page-item active"><a class="page-link" href="#">2</a></li>
                  <li class="page-item"><a class="page-link" href="#">3</a></li>
                  <li class="page-item"><a class="page-link" href="#">4</a></li>
                  <li class="page-item"><a class="page-link" href="#">5</a></li>
                  <li class="page-item">
                    <a class="page-link" href="#">
                      next
                      <!-- Download SVG icon from http://tabler.io/icons/icon/chevron-right -->
                      <svg
                        xmlns="http://www.w3.org/2000/svg"
                        width="24"
                        height="24"
                        viewBox="0 0 24 24"
                        fill="none"
                        stroke="currentColor"
                        stroke-width="2"
                        stroke-linecap="round"
                        stroke-linejoin="round"
                        class="icon icon-1"
                      >
                        <path d="M9 6l6 6l-6 6"></path>
                      </svg>
                    </a>
                  </li>
                </ul>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import c_regimen_res from '@comp/dashboard/c_regimen_res.vue'
import c_sexo_res from '@comp/dashboard/c_sexo_res.vue'
import c_areas_res from '@comp/dashboard/areas.vue'
import { onMounted, ref } from 'vue'
import {
  Regimenes_resumen,
  Resumen_Dashboard,
  Trabajadore_Activos_Area,
  Sexo_cantidad
} from '@wails/services/DashboardService'

interface Regimen {
  Cantidad: number
  Nombre: string
}

const regimenes = ref<Array<Regimen>>([])
const resumen = ref<any>()
const sexos = ref<any>([])
const areas = ref(<any>[])

onMounted(async () => {
  try {
    regimenes.value = await Regimenes_resumen()
    resumen.value = await Resumen_Dashboard()
    areas.value = await Trabajadore_Activos_Area()
    sexos.value = await Sexo_cantidad()
  } catch (error) {
    console.error('Error al cargar los datos:', error)
  }
})

const columns = [
  { field: 'Nombre', title: 'Area' },
  { field: 'Cantidad', title: 'Trabajadores' }
]
</script>

<style lang="scss" scoped>
.page {
  display: grid !important;
  grid-template-rows: min-content min-content min-content;
  height: 100vh;
}

.chart-lg {
  height: 180px;
}

.card-sm {
  margin-bottom: 0.75rem;
}

.card-sm .card-body {
  padding: 0.75rem;
}

.subheader {
  font-size: 0.625rem;
  font-weight: 600;
  text-transform: uppercase;
  letter-spacing: 0.04em;
  line-height: 1.6;
  color: #656d77;
}

.table td,
.table th {
  padding: 0.5rem 0.75rem;
  font-size: 0.875rem;
}

.chart-sparkline {
  width: 4rem;
  height: 1.5rem;
}
</style>
