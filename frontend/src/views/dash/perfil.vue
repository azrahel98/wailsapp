<template>
  <div class="container py-5">
    <!-- Profile Section -->
    <div class="card mb-4">
      <div class="card-body">
        <div class="row align-items-center">
          <!-- Profile Image and Basic Info -->
          <div class="col-md-3 text-center">
            <img
              src="/placeholder.svg?height=180&width=180"
              alt="Profile Photo"
              class="rounded-circle img-fluid mb-3"
              style="width: 180px; height: 180px; object-fit: cover"
            />
            <button class="btn btn-primary btn-sm position-absolute top-0 end-0 m-3">
              EDIT PROFILE
            </button>
          </div>
          <div class="col-md-9">
            <div class="d-flex align-items-center mb-2">
              <span class="badge bg-light text-dark me-2">#5</span>
              <h2 class="mb-0">John Doe</h2>
            </div>
            <p class="text-muted mb-3"><i class="bi bi-geo-alt"></i> San Francisco, CA</p>

            <!-- Achievement Badges -->
            <div class="d-flex gap-3 mb-4">
              <div class="badge bg-success p-2">
                <div class="d-flex align-items-center">
                  <i class="bi bi-trophy me-1"></i>
                  TOP LUXURY HIGH END
                </div>
                <small>2018 - 2023</small>
              </div>
              <div class="badge bg-primary p-2">
                <div class="d-flex align-items-center">
                  <i class="bi bi-star me-1"></i>
                  TOP 1%
                </div>
                <small>2023</small>
              </div>
              <div class="badge bg-info p-2">
                <div class="d-flex align-items-center">
                  <i class="bi bi-award me-1"></i>
                  TOP 10%
                </div>
                <small>2018-2023</small>
              </div>
            </div>

            <!-- Stats -->
            <div class="row">
              <div class="col-6">
                <div class="d-flex align-items-center">
                  <h3 class="mb-0">105</h3>
                  <i class="bi bi-arrow-up-right text-success ms-2"></i>
                  <span class="text-muted ms-2">RENTALS</span>
                </div>
              </div>
              <div class="col-6">
                <div class="d-flex align-items-center">
                  <h3 class="mb-0">105</h3>
                  <i class="bi bi-arrow-down-right text-danger ms-2"></i>
                  <span class="text-muted ms-2">LISTINGS</span>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Contact and Stats Section -->
    <div class="row mb-4">
      <!-- Contact Details -->
      <div class="col-md-6">
        <div class="card h-100">
          <div class="card-body">
            <h5 class="card-title text-uppercase mb-4">Contact Details</h5>
            <div class="mb-3">
              <div class="text-primary">hello@agency.io</div>
              <small class="text-muted">EMAIL ADDRESS</small>
            </div>
            <div class="mb-3">
              <div>512-555-0189</div>
              <small class="text-muted">PHONE NUMBER</small>
            </div>
            <div class="mb-3">
              <div class="d-flex gap-2">
                <a href="#" class="btn btn-sm btn-outline-dark rounded-circle">
                  <i class="bi bi-linkedin"></i>
                </a>
                <a href="#" class="btn btn-sm btn-outline-dark rounded-circle">
                  <i class="bi bi-instagram"></i>
                </a>
                <a href="#" class="btn btn-sm btn-outline-dark rounded-circle">
                  <i class="bi bi-facebook"></i>
                </a>
                <a href="#" class="btn btn-sm btn-outline-dark rounded-circle">
                  <i class="bi bi-twitter"></i>
                </a>
              </div>
              <small class="text-muted">SOCIAL LINKS</small>
            </div>
            <div>
              <div class="d-flex gap-2">
                <img
                  src="/placeholder.svg?height=20&width=20"
                  alt="Spanish"
                  class="rounded-circle"
                />
                <img
                  src="/placeholder.svg?height=20&width=20"
                  alt="English"
                  class="rounded-circle"
                />
                <img
                  src="/placeholder.svg?height=20&width=20"
                  alt="German"
                  class="rounded-circle"
                />
              </div>
              <small class="text-muted">LANGUAGES SPOKEN</small>
            </div>
          </div>
        </div>
      </div>

      <!-- Agent Stats -->
      <div class="col-md-6">
        <div class="card h-100 bg-primary text-white">
          <div class="card-body">
            <h5 class="card-title text-uppercase mb-4">Agent Stats</h5>
            <div class="row text-center">
              <div class="col-4">
                <h2>2</h2>
                <small>PROBLEMS</small>
              </div>
              <div class="col-4">
                <h2>105</h2>
                <small>OFF-MARKET</small>
              </div>
              <div class="col-4">
                <h2>Yes</h2>
                <small>MANAGEMENT</small>
              </div>
            </div>
            <div class="mt-4">
              <div class="d-flex justify-content-between align-items-center mb-2">
                <div>
                  <i class="bi bi-crown me-2"></i>
                  Rent
                </div>
                <small>TOP CATEGORY</small>
              </div>
              <div class="d-flex justify-content-between align-items-center">
                <div>
                  <i class="bi bi-star me-2"></i>
                  Honest
                </div>
                <small>FEEDBACK</small>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Experience Section -->
    <div class="card">
      <div class="card-body">
        <h5 class="card-title text-uppercase mb-4">Experience</h5>
        <div class="timeline">
          <div class="row mb-4">
            <div class="col-2">2021</div>
            <div class="col-10">
              <h6>Sr. Real Estate Agent</h6>
              <p class="mb-0">Agency X (Part Time)</p>
              <small class="text-muted">San Francisco, CA</small>
            </div>
          </div>
          <div class="row mb-4">
            <div class="col-2">2020</div>
            <div class="col-10">
              <h6>Sales Agent</h6>
              <p class="mb-0">Agency X (Part Time)</p>
              <small class="text-muted">San Francisco, CA</small>
            </div>
          </div>
          <div class="row mb-4">
            <div class="col-2">2019</div>
            <div class="col-10">
              <h6>Jr. Sales Agent</h6>
              <p class="mb-0">Agency X (Part Time)</p>
              <small class="text-muted">San Francisco, CA</small>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
// Component logic can be added here if needed
</script>

<style scoped>
.timeline .row:not(:last-child) {
  position: relative;
}

.timeline .row:not(:last-child)::after {
  content: '';
  position: absolute;
  left: 8%;
  top: 30px;
  bottom: -20px;
  width: 2px;
  background-color: #e9ecef;
}

.badge {
  font-weight: normal;
}

.card {
  border: none;
  box-shadow: 0 0 20px rgba(0, 0, 0, 0.08);
}
</style>
