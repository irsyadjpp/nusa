import {
  Box,
  Button,
  Card,
  CardContent,
  Container,
  Typography,
  Grid,
  Chip,
} from '@mui/material';
import {
  Assignment as AssignmentIcon,
  Quiz as QuizIcon,
  School as SchoolIcon,
  AccessTime as TimeIcon,
} from '@mui/icons-material';

const StudentAssessments = () => {
  const assessments = [
    {
      id: '1',
      title: 'Mathematics Chapter 5 Quiz',
      subject: 'Mathematics',
      class: '10-A',
      type: 'Quiz',
      dueDate: '2024-01-20',
      status: 'Pending',
      duration: '30 minutes',
    },
    {
      id: '2',
      title: 'Science Lab Report',
      subject: 'Science',
      class: '10-A',
      type: 'Assignment',
      dueDate: '2024-01-18',
      status: 'Submitted',
      duration: '1 week',
    },
    {
      id: '3',
      title: 'Physics Midterm Exam',
      subject: 'Physics',
      class: '10-A',
      type: 'Exam',
      dueDate: '2024-01-25',
      status: 'Scheduled',
      duration: '2 hours',
    },
    {
      id: '4',
      title: 'English Essay',
      subject: 'English',
      class: '10-A',
      type: 'Assignment',
      dueDate: '2024-01-22',
      status: 'Pending',
      duration: '3 days',
    },
  ];

  const getTypeIcon = (type: string) => {
    switch (type) {
      case 'Quiz':
        return <QuizIcon />;
      case 'Exam':
        return <SchoolIcon />;
      default:
        return <AssignmentIcon />;
    }
  };

  const getStatusColor = (status: string) => {
    switch (status) {
      case 'Submitted':
        return 'success';
      case 'Pending':
        return 'warning';
      case 'Scheduled':
        return 'info';
      case 'Graded':
        return 'primary';
      default:
        return 'default';
    }
  };

  return (
    <Container maxWidth="xl">
      <Box sx={{ mb: 3 }}>
        <Typography variant="h4" component="h1" gutterBottom>
          My Assessments
        </Typography>
        <Typography variant="body2" color="text.secondary">
          View and complete your assessments
        </Typography>
      </Box>

      <Grid container spacing={3}>
        {assessments.map((assessment) => (
          <Grid size={{ xs: 12, md: 6, lg: 4 }} key={assessment.id}>
            <Card variant="outlined">
              <CardContent>
                <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'flex-start', mb: 2 }}>
                  <Box sx={{ display: 'flex', alignItems: 'center', gap: 1 }}>
                    <Box sx={{ color: 'primary.main' }}>
                      {getTypeIcon(assessment.type)}
                    </Box>
                    <Typography variant="subtitle1" fontWeight="medium">
                      {assessment.title}
                    </Typography>
                  </Box>
                  <Chip
                    label={assessment.status}
                    size="small"
                    color={getStatusColor(assessment.status) as any}
                  />
                </Box>
                <Typography variant="body2" color="text.secondary" gutterBottom>
                  {assessment.subject} • Class {assessment.class}
                </Typography>
                <Box sx={{ display: 'flex', alignItems: 'center', gap: 0.5, mb: 1 }}>
                  <TimeIcon fontSize="small" sx={{ color: 'text.secondary' }} />
                  <Typography variant="body2" color="text.secondary">
                    Duration: {assessment.duration}
                  </Typography>
                </Box>
                <Box sx={{ display: 'flex', alignItems: 'center', gap: 0.5, mb: 2 }}>
                  <Typography variant="body2" color="text.secondary">
                    Due: {assessment.dueDate}
                  </Typography>
                </Box>
                <Button
                  variant={assessment.status === 'Pending' ? 'contained' : 'outlined'}
                  fullWidth
                  disabled={assessment.status === 'Submitted'}
                >
                  {assessment.status === 'Submitted' ? 'View Submission' : assessment.status === 'Scheduled' ? 'View Details' : 'Start Assessment'}
                </Button>
              </CardContent>
            </Card>
          </Grid>
        ))}
      </Grid>
    </Container>
  );
};

export default StudentAssessments;
