import {
  Box,
  Card,
  CardContent,
  Container,
  Typography,
  Grid,
  Chip,
} from '@mui/material';
import {
  AccessTime as TimeIcon,
  LocationOn as LocationIcon,
  School as SchoolIcon,
} from '@mui/icons-material';

const StudentSchedule = () => {
  const schedule = [
    {
      day: 'Monday',
      classes: [
        { subject: 'Mathematics', class: '10-A', time: '08:00-09:30', room: 'Room 101', teacher: 'Mr. Smith' },
        { subject: 'Science', class: '10-A', time: '10:00-11:30', room: 'Room 205', teacher: 'Ms. Johnson' },
        { subject: 'Physics', class: '10-A', time: '13:00-14:30', room: 'Lab 301', teacher: 'Dr. Williams' },
      ],
    },
    {
      day: 'Tuesday',
      classes: [
        { subject: 'English', class: '10-A', time: '08:00-09:30', room: 'Room 102', teacher: 'Mrs. Brown' },
        { subject: 'Mathematics', class: '10-A', time: '10:00-11:30', room: 'Room 101', teacher: 'Mr. Smith' },
        { subject: 'Science', class: '10-A', time: '14:00-15:30', room: 'Room 205', teacher: 'Ms. Johnson' },
      ],
    },
    {
      day: 'Wednesday',
      classes: [
        { subject: 'Mathematics', class: '10-A', time: '08:00-09:30', room: 'Room 101', teacher: 'Mr. Smith' },
        { subject: 'Physics', class: '10-A', time: '10:00-11:30', room: 'Lab 301', teacher: 'Dr. Williams' },
        { subject: 'English', class: '10-A', time: '13:00-14:30', room: 'Room 102', teacher: 'Mrs. Brown' },
      ],
    },
    {
      day: 'Thursday',
      classes: [
        { subject: 'Science', class: '10-A', time: '08:00-09:30', room: 'Room 205', teacher: 'Ms. Johnson' },
        { subject: 'Mathematics', class: '10-A', time: '10:00-11:30', room: 'Room 101', teacher: 'Mr. Smith' },
        { subject: 'Physics', class: '10-A', time: '14:00-15:30', room: 'Lab 301', teacher: 'Dr. Williams' },
      ],
    },
    {
      day: 'Friday',
      classes: [
        { subject: 'English', class: '10-A', time: '08:00-09:30', room: 'Room 102', teacher: 'Mrs. Brown' },
        { subject: 'Science', class: '10-A', time: '10:00-11:30', room: 'Room 205', teacher: 'Ms. Johnson' },
        { subject: 'Mathematics', class: '10-A', time: '13:00-14:30', room: 'Room 101', teacher: 'Mr. Smith' },
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
          View your weekly class schedule
        </Typography>
      </Box>

      <Grid container spacing={3}>
        {schedule.map((daySchedule, index) => (
          <Grid size={{ xs: 12, md: 6, lg: 4 }} key={index}>
            <Card sx={{ height: '100%' }}>
              <CardContent>
                <Typography variant="h6" gutterBottom sx={{ color: 'primary.main' }}>
                  {daySchedule.day}
                </Typography>
                {daySchedule.classes.map((classItem, classIndex) => (
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
                      <TimeIcon fontSize="small" sx={{ color: 'text.secondary' }} />
                      <Typography variant="body2" color="text.secondary">
                        {classItem.time}
                      </Typography>
                    </Box>
                    <Box sx={{ display: 'flex', alignItems: 'center', gap: 0.5, mb: 0.5 }}>
                      <LocationIcon fontSize="small" sx={{ color: 'text.secondary' }} />
                      <Typography variant="body2" color="text.secondary">
                        {classItem.room}
                      </Typography>
                    </Box>
                    <Typography variant="body2" color="text.secondary">
                      {classItem.teacher}
                    </Typography>
                    <Chip label={classItem.class} size="small" variant="outlined" sx={{ mt: 1 }} />
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

export default StudentSchedule;
