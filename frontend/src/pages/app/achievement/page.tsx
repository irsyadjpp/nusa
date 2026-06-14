import React, { useState } from 'react';
import {
  Box,
  Typography,
  Button,
  Grid,
  Paper,
  TextField,
  CircularProgress,
  Alert,
  Card,
  Chip,
  Tab,
  Tabs,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  LinearProgress,
  Container,
  Divider,
} from '@mui/material';
import { useNavigate } from 'react-router-dom';
import { useStudentAchievement, useStudentProgress, useClassAchievement } from '@/services/queries/AchievementQueryService';

const AchievementPage: React.FC = () => {
  const navigate = useNavigate();
  const [activeTab, setActiveTab] = useState<'student' | 'class'>('student');
  const [studentId, setStudentId] = useState('');
  const [classId, setClassId] = useState('');
  const [tpId, setTpId] = useState('');
  const [subjectId, setSubjectId] = useState('');
  const [phaseId, setPhaseId] = useState('');
  const [selectedStudent, setSelectedStudent] = useState<string | null>(null);

  const {
    data: studentAchievements,
    isLoading: isLoadingStudent,
    error: studentError,
  } = useStudentAchievement(selectedStudent || '', { tp_id: tpId || undefined }, {
    enabled: !!selectedStudent && activeTab === 'student',
  });

  const {
    data: studentProgress,
    isLoading: isLoadingProgress,
    error: progressError,
  } = useStudentProgress(selectedStudent || '', { subject_id: subjectId || undefined, phase_id: phaseId || undefined }, {
    enabled: !!selectedStudent && activeTab === 'student',
  });

  const {
    data: classAchievement,
    isLoading: isLoadingClass,
    error: classError,
  } = useClassAchievement(classId || '', { subject_id: subjectId || undefined }, {
    enabled: !!classId && activeTab === 'class',
  });

  const handleSearchStudent = () => {
    if (studentId) {
      setSelectedStudent(studentId);
      setActiveTab('student');
    }
  };

  const handleSearchClass = () => {
    if (classId) {
      setActiveTab('class');
    }
  };

  const renderStudentView = () => {
    if (!selectedStudent) {
      return (
        <Alert severity="info">Enter a student ID to view their achievement</Alert>
      );
    }

    if (isLoadingStudent || isLoadingProgress) {
      return <CircularProgress />;
    }

    if (studentError || progressError) {
      return <Alert severity="error">{studentError?.message || progressError?.message || 'Failed to load achievement data'}</Alert>;
    }

    return (
      <Box sx={{ display: 'flex', flexDirection: 'column', gap: 3 }}>
        {/* Filters */}
        <Paper sx={{ p: 2 }}>
          <Grid container spacing={2} alignItems="center">
            <Grid size={{ xs: 12, md: 3 }}>
              <TextField
                label="TP ID"
                value={tpId}
                onChange={(e) => setTpId(e.target.value)}
                fullWidth
                size="small"
              />
            </Grid>
            <Grid size={{ xs: 12, md: 3 }}>
              <TextField
                label="Subject ID"
                value={subjectId}
                onChange={(e) => setSubjectId(e.target.value)}
                fullWidth
                size="small"
              />
            </Grid>
            <Grid size={{ xs: 12, md: 3 }}>
              <TextField
                label="Phase ID"
                value={phaseId}
                onChange={(e) => setPhaseId(e.target.value)}
                fullWidth
                size="small"
              />
            </Grid>
            <Grid size={{ xs: 12, md: 3 }}>
              <Button variant="contained" onClick={handleSearchStudent}>
                Refresh
              </Button>
            </Grid>
          </Grid>
        </Paper>

        {/* Student Achievement */}
        <Paper sx={{ p: 2 }}>
          <Typography variant="h6" gutterBottom>
            Student Achievement
          </Typography>
          {studentAchievements && studentAchievements.length > 0 ? (
            <TableContainer>
              <Table>
                <TableHead>
                  <TableRow>
                    <TableCell>Competency</TableCell>
                    <TableCell>Mastery Level</TableCell>
                    <TableCell>Score</TableCell>
                    <TableCell>Percentage</TableCell>
                  </TableRow>
                </TableHead>
                <TableBody>
                  {studentAchievements.map((achievement: any) => (
                    <TableRow key={achievement.competency_id}>
                      <TableCell>{achievement.competency_name}</TableCell>
                      <TableCell>
                        <Chip
                          label={achievement.mastery_level}
                          color={
                            achievement.mastery_level === 'EXEMPLARY'
                              ? 'success'
                              : achievement.mastery_level === 'PROFICIENT'
                              ? 'primary'
                              : achievement.mastery_level === 'DEVELOPING'
                              ? 'warning'
                              : 'error'
                          }
                          size="small"
                        />
                      </TableCell>
                      <TableCell>
                        {achievement.score} / {achievement.max_score}
                      </TableCell>
                      <TableCell>{achievement.percentage}%</TableCell>
                    </TableRow>
                  ))}
                </TableBody>
              </Table>
            </TableContainer>
          ) : (
            <Typography color="text.secondary">No achievement data available</Typography>
          )}
        </Paper>

        {/* Competency Progress */}
        <Paper sx={{ p: 2 }}>
          <Typography variant="h6" gutterBottom>
            Competency Progress
          </Typography>
          {studentProgress && studentProgress.length > 0 ? (
            <TableContainer>
              <Table>
                <TableHead>
                  <TableRow>
                    <TableCell>Competency</TableCell>
                    <TableCell>Progress</TableCell>
                    <TableCell>Average Score</TableCell>
                    <TableCell>Mastery Level</TableCell>
                  </TableRow>
                </TableHead>
                <TableBody>
                  {studentProgress.map((progress: any) => (
                    <TableRow key={progress.competency_id}>
                      <TableCell>{progress.competency_name}</TableCell>
                      <TableCell>
                        <Box sx={{ display: 'flex', alignItems: 'center', gap: 1 }}>
                          <LinearProgress
                            variant="determinate"
                            value={progress.progress_percentage}
                            sx={{ flex: 1, minWidth: 100 }}
                          />
                          <Typography variant="body2">{progress.progress_percentage}%</Typography>
                        </Box>
                      </TableCell>
                      <TableCell>{progress.average_score}</TableCell>
                      <TableCell>
                        <Chip
                          label={progress.mastery_level}
                          color={
                            progress.mastery_level === 'EXEMPLARY'
                              ? 'success'
                              : progress.mastery_level === 'PROFICIENT'
                              ? 'primary'
                              : progress.mastery_level === 'DEVELOPING'
                              ? 'warning'
                              : 'error'
                          }
                          size="small"
                        />
                      </TableCell>
                    </TableRow>
                  ))}
                </TableBody>
              </Table>
            </TableContainer>
          ) : (
            <Typography color="text.secondary">No progress data available</Typography>
          )}
        </Paper>

        <Button
          variant="outlined"
          onClick={() => navigate('/reports/generate', { state: { studentId: selectedStudent, classId: null } })}
        >
          Generate Narrative Report
        </Button>
      </Box>
    );
  };

  const renderClassView = () => {
    if (!classId) {
      return (
        <Alert severity="info">Enter a class ID to view class achievement</Alert>
      );
    }

    if (isLoadingClass) {
      return <CircularProgress />;
    }

    if (classError) {
      return <Alert severity="error">{classError.message || 'Failed to load class achievement'}</Alert>;
    }

    return (
      <Box sx={{ display: 'flex', flexDirection: 'column', gap: 3 }}>
        {/* Filters */}
        <Paper sx={{ p: 2 }}>
          <Grid container spacing={2} alignItems="center">
            <Grid size={{ xs: 12, md: 4 }}>
              <TextField
                label="Subject ID"
                value={subjectId}
                onChange={(e) => setSubjectId(e.target.value)}
                fullWidth
                size="small"
              />
            </Grid>
            <Grid size={{ xs: 12, md: 8 }}>
              <Button variant="contained" onClick={handleSearchClass}>
                Refresh
              </Button>
            </Grid>
          </Grid>
        </Paper>

        {/* Class Achievement Summary */}
        {classAchievement && (
          <Paper sx={{ p: 2 }}>
            <Typography variant="h6" gutterBottom>
              Class Achievement Summary
            </Typography>
            <Box sx={{ display: 'flex', flexDirection: 'column', gap: 2 }}>
              <Box>
                <Typography variant="body2" color="text.secondary">
                  Total Students
                </Typography>
                <Typography variant="h4">{classAchievement.total_students}</Typography>
              </Box>
              <Box>
                <Typography variant="body2" color="text.secondary">
                  Average Mastery
                </Typography>
                <Typography variant="h4">{classAchievement.average_mastery}%</Typography>
              </Box>
            </Box>
          </Paper>
        )}

        {/* Competency Achievements */}
        {classAchievement?.competency_achievements && classAchievement?.competency_achievements.length > 0 && (
          <Paper sx={{ p: 2 }}>
            <Typography variant="h6" gutterBottom>
              Competency Achievements
            </Typography>
            <TableContainer>
              <Table>
                <TableHead>
                  <TableRow>
                    <TableCell>Competency</TableCell>
                    <TableCell>Average Score</TableCell>
                    <TableCell>Mastery Distribution</TableCell>
                  </TableRow>
                </TableHead>
                <TableBody>
                  {classAchievement?.competency_achievements.map((comp: any) => (
                    <TableRow key={comp.competency_id}>
                      <TableCell>{comp.competency_name}</TableCell>
                      <TableCell>{comp.average_score}</TableCell>
                      <TableCell>
                        <Box sx={{ display: 'flex', gap: 1, flexWrap: 'wrap' }}>
                          <Chip label={`Excellent: ${comp.mastery_distribution.excellent}`} size="small" color="success" />
                          <Chip label={`Proficient: ${comp.mastery_distribution.proficient}`} size="small" color="primary" />
                          <Chip label={`Developing: ${comp.mastery_distribution.developing}`} size="small" color="warning" />
                          <Chip label={`Beginning: ${comp.mastery_distribution.beginning}`} size="small" color="error" />
                        </Box>
                      </TableCell>
                    </TableRow>
                  ))}
                </TableBody>
              </Table>
            </TableContainer>
          </Paper>
        )}

        {/* Areas for Improvement */}
        {classAchievement?.areas_for_improvement && classAchievement?.areas_for_improvement.length > 0 && (
          <Paper sx={{ p: 2 }}>
            <Typography variant="h6" gutterBottom>
              Areas for Improvement
            </Typography>
            <Box sx={{ display: 'flex', flexDirection: 'column', gap: 1 }}>
              {classAchievement?.areas_for_improvement.map((area: any) => (
                <Card key={area.competency_id} sx={{ p: 1 }}>
                  <Typography variant="subtitle2" gutterBottom>
                    {area.competency_name}
                  </Typography>
                  <Typography variant="body2" color="text.secondary">
                    Average Score: {area.average_score}
                  </Typography>
                  <Typography variant="body2" color="text.secondary">
                    Struggling Students: {area.struggling_students}
                  </Typography>
                </Card>
              ))}
            </Box>
          </Paper>
        )}

        {/* Top Performers */}
        {classAchievement?.top_performers && classAchievement?.top_performers.length > 0 && (
          <Paper sx={{ p: 2 }}>
            <Typography variant="h6" gutterBottom>
              Top Performers
            </Typography>
            <TableContainer>
              <Table>
                <TableHead>
                  <TableRow>
                    <TableCell>Student</TableCell>
                    <TableCell>Average Score</TableCell>
                  </TableRow>
                </TableHead>
                <TableBody>
                  {classAchievement?.top_performers.map((student: any) => (
                    <TableRow key={student.student_id}>
                      <TableCell>{student.student_name}</TableCell>
                      <TableCell>{student.average_score}</TableCell>
                    </TableRow>
                  ))}
                </TableBody>
              </Table>
            </TableContainer>
          </Paper>
        )}
      </Box>
    );
  };

  return (
    <Container maxWidth="xl" sx={{ py: 4 }}>
      <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 4 }}>
        <Typography variant="h4" component="h1">
          Achievement Dashboard
        </Typography>
        <Button variant="outlined" onClick={() => navigate('/evaluation')}>
          Back to Evaluation
        </Button>
      </Box>

      <Paper sx={{ mb: 3 }}>
        <Tabs value={activeTab} onChange={(_, newValue) => setActiveTab(newValue as 'student' | 'class')}>
          <Tab label="Student Achievement" value="student" />
          <Tab label="Class Achievement" value="class" />
        </Tabs>
        <Divider />
      </Paper>

      <Paper sx={{ p: 3 }}>
        {activeTab === 'student' ? (
          <Box sx={{ display: 'flex', flexDirection: 'column', gap: 2 }}>
            <Typography variant="h6" gutterBottom>
              Search by Student
            </Typography>
            <Grid container spacing={2} alignItems="center">
              <Grid size={{ xs: 12, md: 8 }}>
                <TextField
                  label="Student ID"
                  value={studentId}
                  onChange={(e) => setStudentId(e.target.value)}
                  fullWidth
                />
              </Grid>
              <Grid size={{ xs: 12, md: 4 }}>
                <Button variant="contained" onClick={handleSearchStudent} fullWidth>
                  Search Student
                </Button>
              </Grid>
            </Grid>
          </Box>
        ) : (
          <Box sx={{ display: 'flex', flexDirection: 'column', gap: 2 }}>
            <Typography variant="h6" gutterBottom>
              Search by Class
            </Typography>
            <Grid container spacing={2} alignItems="center">
              <Grid size={{ xs: 12, md: 8 }}>
                <TextField
                  label="Class ID"
                  value={classId}
                  onChange={(e) => setClassId(e.target.value)}
                  fullWidth
                />
              </Grid>
              <Grid size={{ xs: 12, md: 4 }}>
                <Button variant="contained" onClick={handleSearchClass} fullWidth>
                  Search Class
                </Button>
              </Grid>
            </Grid>
          </Box>
        )}
      </Paper>

      {activeTab === 'student' ? renderStudentView() : renderClassView()}
    </Container>
  );
};

export default AchievementPage;
