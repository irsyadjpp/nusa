import React, { useState } from 'react';
import { Box, Grid, Paper, Divider, Stack, Typography, Button } from '@mui/material';
import { AchievementDashboardHeader, AchievementOverview, StudentAchievementList, StudentDetailPanel, CompetencyProgress, StudentTrajectory } from '@/features/achievement';
import { useStudentAchievement, useClassAchievement, useStudentProgress } from '@/services/queries/AchievementQueryService';

const AchievementPage: React.FC = () => {
  const [selectedStudentId, setSelectedStudentId] = useState<string>('');
  const [selectedClassId] = useState<string>('class-1'); // TODO: Get from context
  const [viewMode, setViewMode] = useState<'overview' | 'student' | 'class'>('overview');

  // Query data
  const { data: classAchievement, isLoading: classLoading, refetch: refetchClass } = useClassAchievement(selectedClassId);
  const { data: studentAchievement, refetch: refetchStudent } = useStudentAchievement(selectedStudentId, { enabled: !!selectedStudentId });
  const { data: studentProgress, refetch: refetchProgress } = useStudentProgress(selectedStudentId, { enabled: !!selectedStudentId });

  const handleRefresh = () => {
    if (viewMode === 'class') {
      refetchClass();
    } else if (viewMode === 'student' && selectedStudentId) {
      refetchStudent();
      refetchProgress();
    }
  };

  const handleSelectStudent = (student: any) => {
    setSelectedStudentId(student.id);
    setViewMode('student');
  };

  const handleViewOverview = () => {
    setViewMode('overview');
    setSelectedStudentId('');
  };

  const handleViewClass = () => {
    setViewMode('class');
    setSelectedStudentId('');
  };

  return (
    <Box sx={{ height: '100%', display: 'flex', flexDirection: 'column' }}>
      <AchievementDashboardHeader
        title="Achievement Dashboard"
        onRefresh={handleRefresh}
      />

      <Grid container spacing={2} sx={{ flex: 1, overflow: 'hidden' }}>
        {/* Left Panel - Student List or Overview */}
        <Grid size={{ xs: 12, md: 4 }} sx={{ height: '100%', overflow: 'auto' }}>
          <Paper sx={{ height: '100%', display: 'flex', flexDirection: 'column' }}>
            <Box sx={{ p: 2 }}>
              <Stack direction="row" spacing={1} sx={{ mb: 2 }}>
                <Button
                  variant={viewMode === 'overview' ? 'contained' : 'outlined'}
                  size="small"
                  onClick={handleViewOverview}
                >
                  Overview
                </Button>
                <Button
                  variant={viewMode === 'class' ? 'contained' : 'outlined'}
                  size="small"
                  onClick={handleViewClass}
                >
                  Class View
                </Button>
              </Stack>
            </Box>
            <Divider />
            <Box sx={{ flex: 1, overflow: 'auto' }}>
              {viewMode === 'overview' && classAchievement && (
                <Box sx={{ p: 2 }}>
                  <AchievementOverview
                    stats={{
                      totalStudents: classAchievement.total_students || 0,
                      totalAssessments: classAchievement.total_students || 0,
                      averageScore: classAchievement.average_mastery || 0,
                      completionRate: 85,
                    }}
                  />
                </Box>
              )}
              {viewMode === 'class' && classAchievement && (
                <Box sx={{ p: 2 }}>
                  <Typography variant="subtitle1" gutterBottom>
                    Students
                  </Typography>
                  <StudentAchievementList
                    students={classAchievement.top_performers || []}
                    onSelect={handleSelectStudent}
                    loading={classLoading}
                  />
                </Box>
              )}
              {viewMode === 'student' && (
                <Box sx={{ p: 2 }}>
                  <Button variant="outlined" onClick={handleViewOverview} fullWidth>
                    Back to Overview
                  </Button>
                </Box>
              )}
            </Box>
          </Paper>
        </Grid>

        {/* Right Panel - Student Detail or Class Metrics */}
        <Grid size={{ xs: 12, md: 8 }} sx={{ height: '100%', overflow: 'auto' }}>
          <Paper sx={{ height: '100%', p: 2 }}>
            {viewMode === 'overview' && classAchievement && (
              <Box>
                <Typography variant="h6" gutterBottom>
                  Class Achievement Overview
                </Typography>
                <AchievementOverview
                  stats={{
                    totalStudents: classAchievement.total_students || 0,
                    totalAssessments: classAchievement.total_students || 0,
                    averageScore: classAchievement.average_mastery || 0,
                    completionRate: 85,
                  }}
                />
              </Box>
            )}

            {viewMode === 'class' && classAchievement && (
              <Box>
                <Typography variant="h6" gutterBottom>
                  Class Metrics
                </Typography>
                <CompetencyProgress competencies={classAchievement.competency_achievements || []} />
              </Box>
            )}

            {viewMode === 'student' && selectedStudentId && (
              <Box>
                <Typography variant="h6" gutterBottom>
                  Student Achievement Detail
                </Typography>
                {studentAchievement && studentAchievement[0] && (
                  <StudentDetailPanel
                    student={{
                      id: studentAchievement[0].student_id,
                      name: studentAchievement[0].student_name,
                      averageScore: studentAchievement[0].percentage || 0,
                      completedAssessments: 1,
                      totalAssessments: 1,
                      recentAchievements: studentAchievement[0].achieved_criteria || [],
                    }}
                  />
                )}
                <Divider sx={{ my: 2 }} />
                {studentProgress && (
                  <CompetencyProgress
                    competencies={studentProgress.map((p: any) => ({
                      name: p.competency_name,
                      progress: p.progress_percentage || 0,
                      level: p.mastery_level || 'N/A',
                    }))}
                  />
                )}
                <Divider sx={{ my: 2 }} />
                {selectedStudentId && (
                  <StudentTrajectory
                    trajectory={[
                      {
                        date: new Date().toISOString(),
                        milestone: 'Current Progress',
                        score: studentAchievement?.[0]?.percentage || 0,
                      },
                    ]}
                  />
                )}
              </Box>
            )}

            {!selectedStudentId && viewMode === 'student' && (
              <Box sx={{ display: 'flex', alignItems: 'center', justifyContent: 'center', height: '100%' }}>
                Select a student to view detailed achievement
              </Box>
            )}
          </Paper>
        </Grid>
      </Grid>
    </Box>
  );
};

export default AchievementPage;
