import React from 'react';
import {
  Box,
  Button,
  Card,
  CardContent,
  Container,
  TextField,
  Typography,
  Alert,
  CircularProgress,
  Grid,
} from '@mui/material';
import {
  ArrowBack as ArrowBackIcon,
  Save as SaveIcon,
} from '@mui/icons-material';
import { useUser } from '@/services/queries/UsersQueryService';
import { useUpdateUser } from '@/services/commands/UsersCommandService';
import { useNavigate, useParams } from 'react-router-dom';
import { Formik, Form, Field, FormikHelpers } from 'formik';
import * as Yup from 'yup';

interface UpdateUserFormValues {
  name: string;
  role_id?: string;
  school_id?: string;
  is_active: boolean;
}

const validationSchema = Yup.object({
  name: Yup.string().required('Name is required'),
});

const EditUser = () => {
  const navigate = useNavigate();
  const { id } = useParams<{ id: string }>();
  const [alert, setAlert] = React.useState<{ show: boolean; message: string; severity: 'success' | 'error' }>({
    show: false,
    message: '',
    severity: 'success',
  });

  const { data: user, isLoading, error } = useUser(id!);

  const updateMutation = useUpdateUser({
    onSuccess: () => {
      setAlert({ show: true, message: 'User updated successfully', severity: 'success' });
      setTimeout(() => navigate(`/dashboard/users/${id}`), 1500);
    },
    onError: (error: any) => {
      setAlert({ show: true, message: error.message || 'Failed to update user', severity: 'error' });
    },
  });

  const handleSubmit = (values: UpdateUserFormValues, helpers: FormikHelpers<UpdateUserFormValues>) => {
    updateMutation.mutate({ id: id!, data: values });
    helpers.setSubmitting(false);
  };

  if (isLoading) {
    return (
      <Box display="flex" justifyContent="center" alignItems="center" minHeight="400px">
        <CircularProgress />
      </Box>
    );
  }

  if (error || !user) {
    return (
      <Container maxWidth="xl">
        <Alert severity="error" sx={{ mt: 3 }}>
          Failed to load user details
        </Alert>
      </Container>
    );
  }

  const initialValues: UpdateUserFormValues = {
    name: user.name,
    role_id: user.role_id || '',
    school_id: user.school_id || '',
    is_active: user.is_active,
  };

  return (
    <Container maxWidth="xl">
      <Box sx={{ mb: 3 }}>
        <Button
          startIcon={<ArrowBackIcon />}
          onClick={() => navigate(`/dashboard/users/${id}`)}
          sx={{ mb: 2 }}
        >
          Back to User Details
        </Button>
        <Typography variant="h4" component="h1" gutterBottom>
          Edit User
        </Typography>
        <Typography variant="body2" color="text.secondary">
          Update the user information
        </Typography>
      </Box>

      {alert.show && (
        <Alert
          severity={alert.severity}
          onClose={() => setAlert({ ...alert, show: false })}
          sx={{ mb: 2 }}
        >
          {alert.message}
        </Alert>
      )}

      <Card>
        <CardContent>
          <Formik
            initialValues={initialValues}
            validationSchema={validationSchema}
            onSubmit={handleSubmit}
            enableReinitialize
          >
            {({ values, errors, touched, isSubmitting, handleChange, handleBlur }) => (
              <Form>
                <Grid container spacing={3}>
                  <Grid size={{ xs: 12, md: 6 }}>
                    <Field
                      as={TextField}
                      name="name"
                      label="Full Name"
                      fullWidth
                      required
                      value={values.name}
                      onChange={handleChange}
                      onBlur={handleBlur}
                      error={touched.name && Boolean(errors.name)}
                      helperText={touched.name && errors.name}
                      sx={{ mb: 2 }}
                    />
                  </Grid>
                  <Grid size={{ xs: 12, md: 6 }}>
                    <Field
                      as={TextField}
                      name="email"
                      label="Email"
                      fullWidth
                      value={user.email}
                      disabled
                      sx={{ mb: 2 }}
                      helperText="Email cannot be changed"
                    />
                  </Grid>
                  <Grid size={{ xs: 12, md: 6 }}>
                    <Field
                      as={TextField}
                      name="role_id"
                      label="Role ID"
                      fullWidth
                      value={values.role_id}
                      onChange={handleChange}
                      onBlur={handleBlur}
                      error={touched.role_id && Boolean(errors.role_id)}
                      helperText={touched.role_id && errors.role_id}
                      sx={{ mb: 2 }}
                    />
                  </Grid>
                  <Grid size={{ xs: 12, md: 6 }}>
                    <Field
                      as={TextField}
                      name="school_id"
                      label="School ID"
                      fullWidth
                      value={values.school_id}
                      onChange={handleChange}
                      onBlur={handleBlur}
                      error={touched.school_id && Boolean(errors.school_id)}
                      helperText={touched.school_id && errors.school_id}
                      sx={{ mb: 2 }}
                    />
                  </Grid>
                  <Grid size={{ xs: 12 }}>
                    <Box sx={{ display: 'flex', gap: 2, justifyContent: 'flex-end', mt: 3 }}>
                      <Button
                        variant="outlined"
                        onClick={() => navigate(`/dashboard/users/${id}`)}
                        disabled={isSubmitting}
                      >
                        Cancel
                      </Button>
                      <Button
                        type="submit"
                        variant="contained"
                        startIcon={isSubmitting ? <CircularProgress size={20} /> : <SaveIcon />}
                        disabled={isSubmitting}
                      >
                        {isSubmitting ? 'Saving...' : 'Save Changes'}
                      </Button>
                    </Box>
                  </Grid>
                </Grid>
              </Form>
            )}
          </Formik>
        </CardContent>
      </Card>
    </Container>
  );
};

export default EditUser;
