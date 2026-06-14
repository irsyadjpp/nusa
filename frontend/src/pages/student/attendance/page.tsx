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
  CalendarToday as CalendarIcon,
  CheckCircle as CheckCircleIcon,
  Cancel as CancelIcon,
  Schedule as LateIcon,
} from '@mui/icons-material';

const StudentAttendance = () => {
  const attendanceRecords = [
    { date: '2024-01-15', day: 'Monday', status: 'present', class: 'Mathematics 10-A' },
    { date: '2024-01-14', day: 'Sunday', status: 'weekend', class: '-' },
    { date: '2024-01-13', day: 'Saturday', status: 'weekend', class: '-' },
    { date: '2024-01-12', day: 'Friday', status: 'present', class: 'Science 10-A' },
    { date: '2024-01-11', day: 'Thursday', status: 'late', class: 'Physics 10-A' },
    { date: '2024-01-10', day: 'Wednesday', status: 'present', class: 'Mathematics 10-A' },
    { date: '2024-01-09', day: 'Tuesday', status: 'absent', class: 'English 10-A' },
    { date: '2024-01-08', day: 'Monday', status: 'present', class: 'Mathematics 10-A' },
  ];

  const attendanceSummary = {
    present: 5,
    absent: 1,
    late: 1,
    weekend: 2,
    total: 9,
    percentage: 71,
  };

  const getStatusIcon = (status: string) => {
    switch (status) {
      case 'present':
        return <CheckCircleIcon color="success" />;
      case 'absent':
        return <CancelIcon color="error" />;
      case 'late':
        return <LateIcon color="warning" />;
      default:
        return null;
    }
  };

  const getStatusColor = (status: string) => {
    switch (status) {
      case 'present':
        return 'success';
      case 'absent':
        return 'error';
      case 'late':
        return 'warning';
      case 'weekend':
        return 'default';
      default:
        return 'default';
    }
  };

  return (
    <Container maxWidth="xl">
      <Box sx={{ mb: 3 }}>
        <Typography variant="h4" component="h1" gutterBottom>
          My Attendance
        </Typography>
        <Typography variant="body2" color="text.secondary">
          View your attendance record across all classes
        </Typography>
      </Box>

      <Grid container spacing={3} sx={{ mb: 3 }}>
        <Grid size={{ xs: 12, sm: 6, md: 3 }}>
          <Card>
            <CardContent>
              <Typography variant="body2" color="text.secondary" gutterBottom>
                Present
              </Typography>
              <Typography variant="h4" color="success.main">
                {attendanceSummary.present}
              </Typography>
            </CardContent>
          </Card>
        </Grid>
        <Grid size={{ xs: 12, sm: 6, md: 3 }}>
          <Card>
            <CardContent>
              <Typography variant="body2" color="text.secondary" gutterBottom>
                Absent
              </Typography>
              <Typography variant="h4" color="error.main">
                {attendanceSummary.absent}
              </Typography>
            </CardContent>
          </Card>
        </Grid>
        <Grid size={{ xs: 12, sm: 6, md: 3 }}>
          <Card>
            <CardContent>
              <Typography variant="body2" color="text.secondary" gutterBottom>
                Late
              </Typography>
              <Typography variant="h4" color="warning.main">
                {attendanceSummary.late}
              </Typography>
            </CardContent>
          </Card>
        </Grid>
        <Grid size={{ xs: 12, sm: 6, md: 3 }}>
          <Card>
            <CardContent>
              <Typography variant="body2" color="text.secondary" gutterBottom>
                Attendance Rate
              </Typography>
              <Typography variant="h4" color="primary.main">
                {attendanceSummary.percentage}%
              </Typography>
            </CardContent>
          </Card>
        </Grid>
      </Grid>

      <Card>
        <CardContent>
          <Typography variant="h6" gutterBottom>
            Attendance History
          </Typography>
          {attendanceRecords.map((record, index) => (
            <Box
              key={index}
              sx={{
                display: 'flex',
                justifyContent: 'space-between',
                alignItems: 'center',
                py: 2,
                borderBottom: index < attendanceRecords.length - 1 ? 1 : 0,
                borderColor: 'divider',
              }}
            >
              <Box sx={{ display: 'flex', alignItems: 'center', gap: 2 }}>
                <CalendarIcon color="action" />
                <Box>
                  <Typography variant="subtitle1" fontWeight="medium">
                    {record.day}, {record.date}
                  </Typography>
                  <Typography variant="body2" color="text.secondary">
                    {record.class}
                  </Typography>
                </Box>
              </Box>
              <Box sx={{ display: 'flex', alignItems: 'center', gap: 1 }}>
                {getStatusIcon(record.status)}
                <Chip
                  label={record.status.charAt(0).toUpperCase() + record.status.slice(1)}
                  size="small"
                  color={getStatusColor(record.status) as any}
                />
              </Box>
            </Box>
          ))}
        </CardContent>
      </Card>
    </Container>
  );
};

export default StudentAttendance;
