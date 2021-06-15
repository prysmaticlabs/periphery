import { TestBed } from '@angular/core/testing';
import { AuthGuard } from './auth.guard';
import { Router } from '@angular/router';
import { AuthService } from '../services/auth.service';

describe('AuthGuard', () => {
  let guard: AuthGuard;
  let authService: AuthService;
  const router = {
    navigate: jasmine.createSpy('navigate')
  };
  const authMock = {};

  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [
        AuthGuard,
        {
          provide: AuthService,
          useValue: authMock,
        },
        {
          provide: Router,
          useValue: router,
        }
      ],
    });
    guard = TestBed.inject(AuthGuard);
    authService = TestBed.inject(AuthService);
    router.navigate.and.callFake(() => {});
  });

  describe('canActivate', () => {
    const routeMock: any = { snapshot: {} };
    const routeStateMock: any = { snapshot: {}, url: '/'};

    it('should return true for a logged in user', () => {
      authService.token = 'hello';
      expect(guard.canActivate(routeMock, routeStateMock)).toEqual(true);
    });

    it('should return false for an unauthenticated user', () => {
      authService.token = '';
      expect(guard.canActivate(routeMock, routeStateMock)).toEqual(false);
    });
  });
});
