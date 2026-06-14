import {
  Box,
  Card,
  CardContent,
  Container,
  Typography,
  Grid,
  Chip,
  CircularProgress,
  Alert,
} from '@mui/material';
import {
  School as SchoolIcon,
  People as PeopleIcon,
  Assignment as AssignmentIcon,
  Event as EventIcon,
  CheckCircle as CheckCircleIcon,
} from '@mui/icons-material';
import { useClasses } from '@/services/queries/ClassesQueryService';
import { useSchedules } from '@/services/queries/SchedulesQueryService';

const TeacherDashboard = () => {
  const { data: classes, isLoading: classesLoading, error: classesError } = useClasses();
  const { data: schedules, isLoading: schedulesLoading, error: schedulesError } = useSchedules();

  if (classesLoading || schedulesLoading) {
    return (
      <Container maxWidth="xl">
        <Box sx={{ display: 'flex', justifyContent: 'center', mt: 10 }}>
          <CircularProgress />
        </Box>
      </Container>
    );
  }

  if (classesError || schedulesError) {
    return (
      <Container maxWidth="xl">
        <Alert severity="error">
          {classesError?.message || schedulesError?.message || 'Failed to load data'}
        </Alert>
      </Container>
    );
  }

  const totalStudents = classes?.data?.reduce((sum: number, cls: any) => sum + (cls.student_count || 0), 0) || 0;
  const todaySchedule = schedules?.data?.filter((schedule: any) => {
    const today = new Date().toISOString().split('T')[0];
    return schedule.date === today;
  }) || [];

  const stats = [
    { title: 'My Classes', value: classes?.data?.length || '0', icon: <SchoolIcon />, color: 'primary' },
    { title: 'Total Students', value: totalStudents, icon: <PeopleIcon />, color: 'success' },
    { title: 'Pending Assignments', value: '3', icon: <AssignmentIcon />, color: 'warning' },
    { title: 'Upcoming Classes', value: todaySchedule.length, icon: <EventIcon />, color: 'info' },
  ];

  const recentActivities = [
    { title: 'Grade Assignment: Mathematics - Chapter 5', time: '2 hours ago', type: 'grading' },
    { title: 'Create Assessment: Science Quiz', time: '5 hours ago', type: 'assessment' },
    { title: 'Mark Attendance: Class 10-A', time: '1 day ago', type: 'attendance' },
    { title: 'Update Student Progress: John Doe', time: '2 days ago', type: 'progress' },
  ];

  const upcomingSchedule = todaySchedule.map((schedule: any) => ({
    subject: schedule.subject || 'Class',
    class: schedule.class_name || schedule.class || 'N/A',
    time: `${schedule.start_time} - ${schedule.end_time}`,
    room: schedule.room || 'TBD',
  })) || [
    { subject: 'Mathematics', class: '10-A', time: '08:00 - 09:30', room: 'Room 101' },
    { subject: 'Science', class: '10-B', time: '10:00 - 11:30', room: 'Room 205' },
    { subject: 'Mathematics', class: '9-A', time: '13:00 - 14:30', room: 'Room 101' },
  ];

  return (
    <Container maxWidth="xl">
      <Box sx={{ mb: 3 }}>
        <Typography variant="h4" component="h1" gutterBottom>
          Teacher Dashboard
        </Typography>
        <Typography variant="body2" color="text.secondary">
          Welcome back! Here's an overview of your teaching activities.
        </Typography>
      </Box>

      <Grid container spacing={3} sx={{ mb: 3 }}>
        {stats.map((stat, index) => (
          <Grid size={{ xs: 12, sm: 6, md: 3 }} key={index}>
            <Card>
              <CardContent>
                <Box sx={{ display: 'flex', alignItems: 'center', justifyContent: 'space-between' }}>
                  <Box>
                    <Typography variant="body2" color="text.secondary" gutterBottom>
                      {stat.title}
                    </Typography>
                    <Typography variant="h4" component="div">
                      {stat.value}
                    </Typography>
                  </Box>
                  <Box sx={{ color: `${stat.color}.main`, fontSize: 40 }}>
                    {stat.icon}
                  </Box>
                </Box>
              </CardContent>
            </Card>
          </Grid>
        ))}
      </Grid>

      <Grid container spacing={3}>
        <Grid size={{ xs: 12, md: 8 }}>
          <Card sx={{ height: '100%' }}>
            <CardContent>
              <Typography variant="h6" gutterBottom>
                Today's Schedule
              </Typography>
              {upcomingSchedule.map((item: any, index: number) => (
                <Box
                  key={index}
                  sx={{
                    display: 'flex',
                    justifyContent: 'space-between',
                    alignItems: 'center',
                    py: 2,
                    borderBottom: index < upcomingSchedule.length - 1 ? 1 : 0,
                    borderColor: 'divider',
                  }}
                >
                  <Box>
                    <Typography variant="subtitle1" fontWeight="medium">
                      {item.subject}
                    </Typography>
                    <Typography variant="body2" color="text.secondary">
                      Class {item.class} • {item.room}
                    </Typography>
                  </Box>
                  <Chip label={item.time} size="small" variant="outlined" />
                </Box>
              ))}
            </CardContent>
          </Card>
        </Grid>

        <Grid size={{ xs: 12, md: 4 }}>
          <Card sx={{ height: '100%' }}>
            <CardContent>
              <Typography variant="h6" gutterBottom>
                Recent Activities
              </Typography>
              {recentActivities.map((activity, index) => (
                <Box
                  key={index}
                  sx={{
                    display: 'flex',
                    alignItems: 'flex-start',
                    py: 2,
                    borderBottom: index < recentActivities.length - 1 ? 1 : 0,
                    borderColor: 'divider',
                  }}
                >
                  <CheckCircleIcon sx={{ mr: 1, color: 'success.main', fontSize: 20, mt: 0.5 }} />
                  <Box>
                    <Typography variant="body2" fontWeight="medium">
                      {activity.title}
                    </Typography>
                    <Typography variant="caption" color="text.secondary">
                      {activity.time}
                    </Typography>
                  </Box>
                </Box>
              ))}
            </CardContent>
          </Card>
        </Grid>
      </Grid>
    </Container>
  );
};

export default TeacherDashboard;
