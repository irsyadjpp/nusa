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
  AccessTime as TimeIcon,
  LocationOn as LocationIcon,
  School as SchoolIcon,
} from '@mui/icons-material';
import { useSchedules } from '@/services/queries/SchedulesQueryService';

const TeacherSchedule = () => {
  const { data: schedulesData, isLoading, error } = useSchedules();

  if (isLoading) {
    return (
      <Container maxWidth="xl">
        <Box sx={{ display: 'flex', justifyContent: 'center', mt: 10 }}>
          <CircularProgress />
        </Box>
      </Container>
    );
  }

  if (error) {
    return (
      <Container maxWidth="xl">
        <Alert severity="error">{error.message}</Alert>
      </Container>
    );
  }

  const schedule = schedulesData?.data || [
    {
      day: 'Monday',
      classes: [
        { subject: 'Mathematics', class: '10-A', time: '08:00-09:30', room: 'Room 101' },
        { subject: 'Mathematics', class: '9-A', time: '10:00-11:30', room: 'Room 101' },
        { subject: 'Mathematics', class: '10-A', time: '13:00-14:30', room: 'Room 101' },
      ],
    },
    {
      day: 'Tuesday',
      classes: [
        { subject: 'Science', class: '10-B', time: '10:00-11:30', room: 'Room 205' },
        { subject: 'Physics', class: '11-A', time: '14:00-15:30', room: 'Lab 301' },
      ],
    },
    {
      day: 'Wednesday',
      classes: [
        { subject: 'Mathematics', class: '10-A', time: '08:00-09:30', room: 'Room 101' },
        { subject: 'Mathematics', class: '9-A', time: '10:00-11:30', room: 'Room 101' },
        { subject: 'Mathematics', class: '10-A', time: '13:00-14:30', room: 'Room 101' },
      ],
    },
    {
      day: 'Thursday',
      classes: [
        { subject: 'Science', class: '10-B', time: '10:00-11:30', room: 'Room 205' },
        { subject: 'Physics', class: '11-A', time: '14:00-15:30', room: 'Lab 301' },
      ],
    },
    {
      day: 'Friday',
      classes: [
        { subject: 'Mathematics', class: '10-A', time: '08:00-09:30', room: 'Room 101' },
        { subject: 'Mathematics', class: '9-A', time: '10:00-11:30', room: 'Room 101' },
        { subject: 'Mathematics', class: '10-A', time: '13:00-14:30', room: 'Room 101' },
      ],
    },
  ];

  return (
    <Container maxWidth="xl">
      <Box sx={{ mb: 3 }}>
        <Typography variant="h4" component="h1" gutterBottom>
          My Schedule
        </Typography>
        <Typography variant="body2" color="text.secondary">
          View your weekly teaching schedule
        </Typography>
      </Box>

      <Grid container spacing={3}>
        {schedule.map((daySchedule: any, index: number) => (
          <Grid size={{ xs: 12, md: 6, lg: 4 }} key={index}>
            <Card sx={{ height: '100%' }}>
              <CardContent>
                <Typography variant="h6" gutterBottom sx={{ color: 'primary.main' }}>
                  {daySchedule.day}
                </Typography>
                {daySchedule.classes.map((classItem: any, classIndex: number) => (
                  <Box
                    key={classIndex}
                    sx={{
                      p: 2,
                      mb: classIndex < daySchedule.classes.length - 1 ? 2 : 0,
                      border: 1,
                      borderColor: 'divider',
                      borderRadius: 1,
                    }}
                  >
                    <Box sx={{ display: 'flex', alignItems: 'center', gap: 0.5, mb: 1 }}>
                      <SchoolIcon fontSize="small" />
                      <Typography variant="subtitle2" fontWeight="medium">
                        {classItem.subject}
                      </Typography>
                    </Box>
                    <Box sx={{ display: 'flex', alignItems: 'center', gap: 0.5, mb: 0.5 }}>
                      <TimeIcon fontSize="small" />
                      <Typography variant="body2" color="text.secondary">
                        {classItem.time}
                      </Typography>
                    </Box>
                    <Box sx={{ display: 'flex', alignItems: 'center', gap: 0.5, mb: 1 }}>
                      <LocationIcon fontSize="small" />
                      <Typography variant="body2" color="text.secondary">
                        {classItem.room}
                      </Typography>
                    </Box>
                    <Chip label={classItem.class} size="small" variant="outlined" />
                  </Box>
                ))}
              </CardContent>
            </Card>
          </Grid>
        ))}
      </Grid>
    </Container>
  );
};

export default TeacherSchedule;
