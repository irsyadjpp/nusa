# UI/UX Specification

**Version:** 1.0  
**Date:** June 13, 2026  
**Based On:** UI_INVENTORY.md, API_SPEC.md, TARGET_ARCHITECTURE.md

---

## Executive Summary

This document specifies the UI/UX design for the NUSA Platform, covering three user portals: Admin Portal, Teacher Portal, and Student Portal. The design prioritizes maintainability, scalability, simplicity, testability, security, and observability while providing a consistent user experience across all portals.

**Design Principles:**
- Consistent navigation and layout
- Role-based access control
- Responsive design for all devices
- Accessibility compliance (WCAG 2.1 AA)
- Progressive enhancement
- Mobile-first approach

---

## Admin Portal

### Navigation Structure

**Top Navigation:**
- Logo / Brand
- School Selector (if multi-school admin)
- User Profile Dropdown
- Notifications Bell
- Settings

**Sidebar Navigation:**
- Dashboard
- Schools
- Users
- Roles & Permissions
- Academic Foundation
  - Academic Years
  - Semesters
  - Subject Categories
  - Graduate Profile Dimensions
  - CP Alignments
  - System Configurations
- Curriculum
  - Subjects
  - Phases
  - Elements
  - Subelements
  - CP Management
- Reports
  - System Reports
  - Audit Logs
- Settings
  - School Settings
  - System Settings

### Screen Inventory

#### 1. Dashboard

**Purpose:** Overview of system status and key metrics

**Components:**
- Statistics Cards
  - Total Schools
  - Total Users
  - Total Classes
  - Active Academic Years
- Charts
  - User Growth Chart
  - Class Distribution Chart
  - System Activity Chart
- Recent Activity Feed
- Quick Actions
  - Add School
  - Add User
  - Create Academic Year

**Empty State:**
- Message: "Welcome to NUSA Platform Admin"
- Illustration: Empty dashboard icon
- Call to Action: "Get started by adding your first school"

**Error State:**
- Message: "Failed to load dashboard data"
- Retry Button
- Support Link

---

#### 2. Schools Management

**Purpose:** Manage schools in the system

**Screens:**
- Schools List
- School Details
- Create/Edit School

**Schools List:**
- Table with columns: Name, Code, Address, Phone, Status, Actions
- Search bar
- Filter by status
- Pagination
- Actions: View, Edit, Delete, Activate/Deactivate

**School Details:**
- School information card
- Statistics (users, classes, academic years)
- Tabs: Overview, Users, Classes, Settings

**Create/Edit School Form:**
- Fields:
  - Name (required, text, max 255)
  - Code (required, text, unique, max 50)
  - Address (optional, textarea)
  - Phone (optional, text)
  - Email (optional, email)
  - Status (required, select: ACTIVE, ARCHIVED)
- Validation: Real-time validation
- Submit Button
- Cancel Button

**Empty State:**
- Message: "No schools found"
- Call to Action: "Add your first school"

---

#### 3. Users Management

**Purpose:** Manage users across all schools

**Screens:**
- Users List
- User Details
- Create/Edit User

**Users List:**
- Table with columns: Name, Email, Role, School, Status, Actions
- Search bar
- Filter by role, school, status
- Pagination
- Actions: View, Edit, Delete, Activate/Deactivate

**User Details:**
- User information card
- Role assignment
- School assignment
- Activity history
- Tabs: Overview, Roles, Activity

**Create/Edit User Form:**
- Fields:
  - Name (required, text, max 255)
  - Email (required, email, unique)
  - Password (required for create, optional for edit, min 8 characters)
  - Role (required, select)
  - School (required, select)
  - Status (required, select: ACTIVE, INACTIVE)
- Validation: Real-time validation
- Submit Button
- Cancel Button

**Empty State:**
- Message: "No users found"
- Call to Action: "Add your first user"

---

#### 4. Academic Foundation

**Purpose:** Manage academic years, semesters, subject categories, etc.

**Screens:**
- Academic Years List
- Academic Year Details
- Semesters List
- Subject Categories List
- Graduate Profile Dimensions List
- CP Alignments List
- System Configurations List

**Academic Years List:**
- Table with columns: Year, School, Start Date, End Date, Status, Actions
- Filter by school, status
- Actions: View, Edit, Activate, Archive
- Status indicators (DRAFT, ACTIVE, ARCHIVED)

**Semesters List:**
- Table with columns: Name, Type, Start Date, End Date, Order, Status, Actions
- Filter by academic year, status
- Actions: View, Edit, Delete

**Subject Categories List:**
- Table with columns: Code, Name, Description, Mandatory, Status, Actions
- Filter by status
- Actions: View, Edit, Delete

**Graduate Profile Dimensions List:**
- Table with columns: Code, Name, Sequence, Description, Status, Actions
- Filter by status
- Actions: View, Edit, Delete

**CP Alignments List:**
- Table with columns: Subject, Dimension, Description, Status, Actions
- Filter by subject, dimension, status
- Actions: View, Edit, Delete
- Bulk Create Button
- Report Button

**System Configurations List:**
- Table with columns: Key, Value, Type, Category, Status, Actions
- Filter by category, status
- Actions: View, Edit, Delete
- System configurations protected (cannot delete)

---

### Page Specifications

#### Academic Year Create/Edit Form

**Form Layout:**
- Two-column layout
- Left column: Basic information
- Right column: Date ranges

**Fields:**
- School (required, select)
- Year (required, text, format: YYYY/YYYY)
- Start Date (required, date picker)
- End Date (required, date picker)
- Status (required, select: DRAFT, ACTIVE, ARCHIVED)

**Validation Rules:**
- School must be active
- Year must be unique per school
- End date must be after start date
- Date ranges cannot overlap with existing academic years
- Only DRAFT academic years can be modified
- Only one ACTIVE academic year per school

**Error States:**
- Field-level errors (red border, error message below field)
- Form-level errors (alert banner at top)
- Conflict errors (specific message about overlap)

**Success State:**
- Success notification
- Redirect to academic years list

---

## Teacher Portal

### Navigation Structure

**Top Navigation:**
- Logo / Brand
- School Selector
- User Profile Dropdown
- Notifications Bell
- Messages Icon

**Sidebar Navigation:**
- Dashboard
- My Classes
- Curriculum
  - CP Management
  - TP Workspace
  - ATP Workspace
  - Modul Ajar
- Assessment
  - Assessments
  - Rubrics
  - Evidence
  - Evaluations
- Students
  - Attendance
  - Grades
  - Achievement
- Schedule
  - My Schedule
  - Timetable
- Reports
  - Narrative Reports
  - Class Reports
- Communication
  - Messages
  - Announcements

### Screen Inventory

#### 1. Dashboard

**Purpose:** Teacher's daily overview

**Components:**
- Today's Schedule Card
  - Upcoming classes with time, room, subject
  - Quick action: Take Attendance
- Statistics Cards
  - Total Students
  - Total Classes
  - Pending Assessments
  - Pending Evaluations
- Upcoming Deadlines
  - Assignments due soon
  - Exams scheduled
- Recent Notifications
- Quick Actions
  - Create Assessment
  - Take Attendance
  - View Schedule

**Empty State:**
- Message: "Welcome to your dashboard"
- Illustration: Empty dashboard icon
- Call to Action: "Get started by setting up your classes"

---

#### 2. My Classes

**Purpose:** Manage teacher's assigned classes

**Screens:**
- Classes List
- Class Details
- Class Students

**Classes List:**
- Card layout or table
- Each class shows: Name, Subject, Grade, Room, Student Count, Schedule
- Filter by academic year, semester, subject
- Actions: View, Manage Students, View Schedule

**Class Details:**
- Class information card
- Schedule card
- Student count
- Tabs: Overview, Students, Schedule, Assessments

**Class Students:**
- Table with columns: Name, Email, Enrollment Date, Status, Actions
- Search bar
- Filter by status
- Actions: View Profile, Remove
- Enroll Student Button

**Empty State:**
- Message: "No classes assigned"
- Call to Action: "Contact administrator to assign classes"

---

#### 3. Attendance

**Purpose:** Record and view student attendance

**Screens:**
- Attendance Recording
- Attendance History
- Attendance Reports

**Attendance Recording:**
- Class selector
- Date picker (defaults to today)
- Student list with attendance status buttons
  - Present (green)
  - Absent (red)
  - Late (yellow)
  - Excused (blue)
- Notes field per student
- Save Button
- Bulk actions: Mark all Present, Mark all Absent

**Attendance History:**
- Table with columns: Date, Class, Present, Absent, Late, Excused, Rate
- Filter by class, date range, student
- Export Button

**Attendance Reports:**
- Class attendance summary
- Student attendance summary
- Charts: Attendance trends, absence reasons
- Export Button

**Empty State:**
- Message: "No attendance records found"
- Call to Action: "Record attendance for today"

---

#### 4. Schedule

**Purpose:** View teacher's schedule and timetable

**Screens:**
- My Schedule
- Timetable
- Calendar View

**My Schedule:**
- List view of today's classes
- Time, Subject, Room, Class
- Quick actions: Take Attendance, View Class

**Timetable:**
- Weekly calendar view
- Days as columns
- Time slots as rows
- Class cards in grid
- Color-coded by subject
- Room information

**Calendar View:**
- Monthly calendar
- Classes marked on dates
- Click for details
- Filter by subject

**Empty State:**
- Message: "No classes scheduled"
- Call to Action: "Contact administrator to assign classes"

---

#### 5. Assessment Workspace

**Purpose:** Create and manage assessments

**Screens:**
- Assessments List
- Assessment Editor
- Rubric Editor
- Evidence Management
- Evaluation Workspace

**Assessments List:**
- Table with columns: Title, Type, Class, Status, Created At, Actions
- Filter by class, type, status
- Actions: View, Edit, Delete, Approve
- Create Assessment Button

**Assessment Editor:**
- Form with sections:
  - Basic Information (Title, Type, Class)
  - Assessment Items (Questions)
  - Answer Key
  - Scoring Guidelines
- AI Generate Button (optional)
- Preview Button
- Save Draft Button
- Submit for Approval Button

**Rubric Editor:**
- Form with sections:
  - Basic Information (Title, Type)
  - Criteria (Name, Description, Weight)
  - Levels (Level, Description, Points)
- Add Criteria Button
- Add Level Button
- Save Button

**Evidence Management:**
- List of student submissions
- Upload area (drag and drop)
- File type icons
- Status indicators
- Actions: View, Download, Evaluate

**Evaluation Workspace:**
- Student selector
- Evidence display
- Rubric-based evaluation form
- Score calculation
- Feedback text area
- Save Button
- Submit Button

**Empty State:**
- Message: "No assessments found"
- Call to Action: "Create your first assessment"

---

### Page Specifications

#### Attendance Recording Form

**Form Layout:**
- Header: Class name, date
- Student list with rows
- Each row: Student name, attendance status buttons, notes

**Fields:**
- Class (required, select)
- Date (required, date picker, default today)
- Attendance Status (required per student, radio buttons)
- Notes (optional per student, textarea)

**Validation Rules:**
- Class must be active
- Date must be valid
- Attendance status required per student
- Cannot record attendance for future dates
- Cannot record attendance if already recorded

**Error States:**
- Field-level errors
- Conflict error (attendance already recorded)
- Class inactive error

**Success State:**
- Success notification
- Update statistics
- Redirect to attendance history

---

## Student Portal

### Navigation Structure

**Top Navigation:**
- Logo / Brand
- User Profile Dropdown
- Notifications Bell
- Messages Icon

**Sidebar Navigation:**
- Dashboard
- My Classes
- Curriculum
  - CP View
  - TP View
  - ATP View
  - Modul Ajar View
- Assessment
  - My Assessments
  - My Submissions
  - My Results
- Attendance
  - My Attendance
  - Attendance History
- Schedule
  - My Schedule
  - Timetable
- Reports
  - My Achievement
  - Narrative Reports
- Communication
  - Messages
  - Announcements

### Screen Inventory

#### 1. Dashboard

**Purpose:** Student's daily overview

**Components:**
- Today's Schedule Card
  - Upcoming classes with time, room, subject
- Statistics Cards
  - Attendance Rate
  - Average Grade
  - Pending Assignments
  - Upcoming Exams
- Upcoming Deadlines
  - Assignments due soon
  - Exams scheduled
- Recent Notifications
- Recent Announcements

**Empty State:**
- Message: "Welcome to your dashboard"
- Illustration: Empty dashboard icon
- Call to Action: "Get started by viewing your schedule"

---

#### 2. My Classes

**Purpose:** View enrolled classes

**Screens:**
- Classes List
- Class Details

**Classes List:**
- Card layout
- Each class shows: Name, Subject, Grade, Teacher, Room, Schedule
- Filter by academic year, semester
- Actions: View Details, View Schedule

**Class Details:**
- Class information card
- Teacher information
- Schedule card
- Tabs: Overview, Schedule, Assessments, Resources

**Empty State:**
- Message: "No classes enrolled"
- Call to Action: "Contact administrator to enroll in classes"

---

#### 3. My Attendance

**Purpose:** View attendance history

**Screens:**
- Attendance Summary
- Attendance History

**Attendance Summary:**
- Statistics cards
  - Total Days
  - Present Days
  - Absent Days
  - Late Days
  - Attendance Rate
- Chart: Attendance trends over time

**Attendance History:**
- Table with columns: Date, Class, Status, Notes
- Filter by class, date range, status
- Export Button

**Empty State:**
- Message: "No attendance records found"
- Call to Action: "Contact teacher to record attendance"

---

#### 4. My Assessments

**Purpose:** View and submit assessments

**Screens:**
- Assessments List
- Assessment Details
- Submission Form

**Assessments List:**
- Table with columns: Title, Type, Class, Due Date, Status, Actions
- Filter by class, status
- Status indicators: Not Started, In Progress, Submitted, Graded
- Actions: View, Submit

**Assessment Details:**
- Assessment information
- Due date countdown
- Instructions
- Submission status
- Submit Button (if not submitted)

**Submission Form:**
- File upload area (drag and drop)
- Supported file types
- File size limit
- Progress indicator
- Submit Button
- Cancel Button

**Empty State:**
- Message: "No assessments assigned"
- Call to Action: "Contact teacher to assign assessments"

---

#### 5. My Results

**Purpose:** View assessment results and grades

**Screens:**
- Results List
- Result Details
- Grade Book

**Results List:**
- Table with columns: Assessment, Class, Score, Grade, Date, Actions
- Filter by class, date range
- Actions: View Details

**Result Details:**
- Score and grade
- Rubric breakdown
- Teacher feedback
- Evidence (if applicable)

**Grade Book:**
- Table with columns: Class, Assessment, Score, Grade
- Filter by class, academic year
- Calculate average
- Export Button

**Empty State:**
- Message: "No results available"
- Call to Action: "Complete assessments to see results"

---

### Page Specifications

#### Assessment Submission Form

**Form Layout:**
- Header: Assessment title, due date
- Instructions section
- File upload area
- Submit Button

**Fields:**
- File (required, file upload)
  - Supported types: PDF, DOC, DOCX, JPG, PNG
  - Max size: 10MB
- Notes (optional, textarea)

**Validation Rules:**
- File required
- File type must be supported
- File size must be within limit
- Cannot submit after due date
- Cannot resubmit if already graded

**Error States:**
- File type error
- File size error
- Due date passed error
- Already submitted error

**Success State:**
- Success notification
- Update submission status
- Redirect to assessments list

---

## Common Components

### Navigation Components

**Sidebar Navigation**
- Collapsible
- Active state indicator
- Icon + label
- Badge for notifications
- Role-based visibility

**Top Navigation**
- Logo/Brand
- User dropdown
- Notifications bell with badge
- Messages icon with badge
- Settings link

### Form Components

**Form Fields**
- Text input (with validation)
- Number input
- Email input
- Password input (with show/hide)
- Date picker
- Time picker
- Select dropdown
- Multi-select
- Textarea
- File upload (drag and drop)
- Checkbox
- Radio buttons
- Toggle switch

**Form Validation**
- Real-time validation
- Error messages below fields
- Form-level error banner
- Success notification on submit

### Table Components

**Data Table**
- Sortable columns
- Filterable columns
- Pagination
- Row actions menu
- Selectable rows
- Bulk actions
- Export button

**Empty State**
- Illustration
- Message
- Call to action button

**Loading State**
- Skeleton loader
- Spinner
- Progress indicator

### Feedback Components

**Notifications**
- Toast notifications (success, error, warning, info)
- Auto-dismiss after 5 seconds
- Manual dismiss option
- Action buttons

**Modals**
- Confirmation dialogs
- Form modals
- Detail modals
- Backdrop overlay
- Close button

**Alerts**
- Info alerts
- Warning alerts
- Error alerts
- Success alerts
- Dismissible

### Display Components

**Cards**
- Information cards
- Statistics cards
- Action cards
- Image cards

**Charts**
- Line charts
- Bar charts
- Pie charts
- Donut charts
- Area charts

**Badges**
- Status badges
- Role badges
- Priority badges

**Progress**
- Progress bars
- Circular progress
- Step progress

---

## Validation Rules

### Common Validation

**Text Fields:**
- Required: Cannot be empty
- Min length: Minimum character count
- Max length: Maximum character count
- Pattern: Regex pattern match

**Email Fields:**
- Required: Cannot be empty
- Format: Valid email format
- Unique: Must be unique in database

**Date Fields:**
- Required: Cannot be empty
- Format: Valid date format
- Range: Must be within valid range
- Future/past: Must be in future/past

**Number Fields:**
- Required: Cannot be empty
- Min: Minimum value
- Max: Maximum value
- Integer: Must be integer

**Select Fields:**
- Required: Must select option
- Options: Must select from provided options

### Specific Validation

**Academic Year:**
- Year format: YYYY/YYYY
- Date range: End date must be after start date
- Overlap: Cannot overlap with existing academic years
- Active: Only one active per school

**Semester:**
- Sequence: Must be unique within academic year
- Date range: Cannot overlap within academic year
- Order: Must be sequential

**Class:**
- Capacity: Max students must be positive
- Teacher: Must have teacher role
- Active: Cannot activate if deleted

**Attendance:**
- Date: Cannot be future date
- Duplicate: Cannot record same date twice
- Status: Must be valid status

**Schedule:**
- Day: Must be 1-7
- Time: End time must be after start time
- Conflict: Cannot overlap with existing schedule

---

## Empty States

### Empty State Pattern

**Components:**
- Illustration (icon or graphic)
- Title (descriptive message)
- Description (optional context)
- Call to Action (primary button)
- Secondary Action (optional link)

**Examples:**

**No Data:**
- Illustration: Empty box icon
- Title: "No items found"
- Description: "There are no items to display"
- Call to Action: "Create your first item"

**No Results:**
- Illustration: Search icon
- Title: "No results found"
- Description: "Try adjusting your search or filters"
- Call to Action: "Clear filters"

**No Access:**
- Illustration: Lock icon
- Title: "You don't have access"
- Description: "Contact your administrator for access"
- Call to Action: "Request access"

---

## Error States

### Error State Pattern

**Components:**
- Error illustration
- Error title
- Error message
- Retry button
- Support link

**Examples:**

**Load Error:**
- Illustration: Error icon
- Title: "Failed to load data"
- Message: "An error occurred while loading the data"
- Call to Action: "Retry"

**Network Error:**
- Illustration: Network icon
- Title: "Network error"
- Message: "Please check your internet connection"
- Call to Action: "Retry"

**Permission Error:**
- Illustration: Lock icon
- Title: "Access denied"
- Message: "You don't have permission to access this resource"
- Call to Action: "Contact support"

---

## Responsive Design

### Breakpoints

- Mobile: < 768px
- Tablet: 768px - 1024px
- Desktop: > 1024px

### Mobile Adaptations

**Navigation:**
- Sidebar becomes bottom navigation or hamburger menu
- Top navigation simplifies
- Tables become cards

**Forms:**
- Single column layout
- Full-width inputs
- Larger touch targets

**Tables:**
- Card view instead of table
- Horizontal scroll for complex tables
- Simplified columns

---

## Accessibility

### WCAG 2.1 AA Compliance

**Keyboard Navigation:**
- All interactive elements keyboard accessible
- Tab order logical
- Focus indicators visible
- Skip to content link

**Screen Readers:**
- ARIA labels for all interactive elements
- Alt text for images
- Semantic HTML
- Live regions for dynamic content

**Color Contrast:**
- Minimum 4.5:1 ratio for normal text
- Minimum 3:1 ratio for large text
- Color not used as only indicator

**Focus Management:**
- Focus trap in modals
- Focus return after modal close
- Focus visible indicators

---

## Conclusion

The UI/UX specification provides comprehensive design guidance for the Admin Portal, Teacher Portal, and Student Portal. The design prioritizes consistency, accessibility, and user experience while maintaining simplicity and avoiding overengineering.

The specification includes detailed screen inventories, page specifications, validation rules, empty states, and error states for all three portals, ensuring a cohesive user experience across the NUSA Platform.
