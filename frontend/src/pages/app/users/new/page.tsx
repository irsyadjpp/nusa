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
import { useCreateUser } from '@/services/commands/UsersCommandService';
import { useNavigate } from 'react-router-dom';
import { Formik, Form, Field, FormikHelpers } from 'formik';
import * as Yup from 'yup';

interface CreateUserFormValues {
  email: string;
  name: string;
  password: string;
  role_id?: string;
  school_id?: string;
}

const validationSchema = Yup.object({
  email: Yup.string().email('Invalid email address').required('Email is required'),
  name: Yup.string().required('Name is required'),
  password: Yup.string().min(8, 'Password must be at least 8 characters').required('Password is required'),
});

const CreateUser = () => {
  const navigate = useNavigate();
  const [alert, setAlert] = React.useState<{ show: boolean; message: string; severity: 'success' | 'error' }>({
    show: false,
    message: '',
    severity: 'success',
  });

  const createMutation = useCreateUser({
    onSuccess: () => {
      setAlert({ show: true, message: 'User created successfully', severity: 'success' });
      setTimeout(() => navigate('/dashboard/users'), 1500);
    },
    onError: (error: any) => {
      setAlert({ show: true, message: error.message || 'Failed to create user', severity: 'error' });
    },
  });

  const initialValues: CreateUserFormValues = {
    email: '',
    name: '',
    password: '',
    role_id: '',
    school_id: '',
  };

  const handleSubmit = (values: CreateUserFormValues, helpers: FormikHelpers<CreateUserFormValues>) => {
    createMutation.mutate(values);
    helpers.setSubmitting(false);
  };

  return (
    <Container maxWidth="xl">
      <Box sx={{ mb: 3 }}>
        <Button
          startIcon={<ArrowBackIcon />}
          onClick={() => navigate('/dashboard/users')}
          sx={{ mb: 2 }}
        >
          Back to Users
        </Button>
        <Typography variant="h4" component="h1" gutterBottom>
          Create New User
        </Typography>
        <Typography variant="body2" color="text.secondary">
          Fill in the details to create a new user
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
                      required
                      type="email"
                      value={values.email}
                      onChange={handleChange}
                      onBlur={handleBlur}
                      error={touched.email && Boolean(errors.email)}
                      helperText={touched.email && errors.email}
                      sx={{ mb: 2 }}
                    />
                  </Grid>
                  <Grid size={{ xs: 12, md: 6 }}>
                    <Field
                      as={TextField}
                      name="password"
                      label="Password"
                      fullWidth
                      required
                      type="password"
                      value={values.password}
                      onChange={handleChange}
                      onBlur={handleBlur}
                      error={touched.password && Boolean(errors.password)}
                      helperText={touched.password && errors.password}
                      sx={{ mb: 2 }}
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
                  <Grid size={{ xs: 12 }}>
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
                        onClick={() => navigate('/dashboard/users')}
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
                        {isSubmitting ? 'Creating...' : 'Create User'}
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

export default CreateUser;
