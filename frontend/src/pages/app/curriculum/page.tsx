/**
 * Curriculum Landing Page
 * Redirects to subjects or shows curriculum overview
 */

import { useEffect } from 'react';
import { useNavigate } from 'react-router-dom';

const CurriculumPage: React.FC = () => {
  const navigate = useNavigate();

  useEffect(() => {
    // Redirect to subjects page for now
    navigate('/dashboard/curriculum/subjects');
  }, [navigate]);

  return null;
};

export default CurriculumPage;
